package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	entity "github.com/resistance22/university_project/Entity"
	utils "github.com/resistance22/university_project/Utils"
	validator "github.com/resistance22/university_project/Validator"
)

type IUserUseCase interface {
	Register(ctx context.Context, user *validator.RegisterBody) (*entity.User, error)
	Login(ctx context.Context, user *validator.LoginBody) (string, error)
}

type userController struct {
	UseCases IUserUseCase
}

func (controller *userController) Register(c *gin.Context) {
	var body validator.RegisterBody
	if err := c.ShouldBindJSON(&body); err != nil {
		httpError := utils.NewHttpError(err.Error(), http.StatusBadRequest, "Invalid Body")
		response := utils.MakeError(httpError)
		c.JSON(httpError.Status, response)
		return
	}

	res, err := controller.UseCases.Register(c, &body)

	if err != nil {
		httpError := utils.NewHttpError(err.Error(), http.StatusBadRequest, "Something Went Wrong!")
		response := utils.MakeError(httpError)
		c.JSON(httpError.Status, response)
		return
	}

	response, err := utils.StructToMapWithJSONKeys(res)

	if err != nil {
		httpError := utils.NewHttpError(err.Error(), http.StatusBadRequest, "Something Went Wrong!")
		response := utils.MakeError(httpError)
		c.JSON(httpError.Status, response)
		return
	}
	delete(response, "password")

	httpResponse := utils.NewHttpResponse(response, http.StatusCreated)
	jsonResponse := utils.MakeResponse(httpResponse, "New User Created")
	c.JSON(httpResponse.Status, jsonResponse)

}

func (controller *userController) Login(c *gin.Context) {
	var body validator.LoginBody
	if err := c.ShouldBindJSON(&body); err != nil {
		httpError := utils.NewHttpError(err.Error(), http.StatusBadRequest, "Invalid Body")
		response := utils.MakeError(httpError)
		c.JSON(httpError.Status, response)
		return
	}

	token, err := controller.UseCases.Login(c, &body)

	if err != nil {
		httpError := utils.NewHttpError(err.Error(), http.StatusUnauthorized, "You Are Not Authorized")
		response := utils.MakeError(httpError)
		c.JSON(httpError.Status, response)
		return
	}

	httpResponse := utils.NewHttpResponse(gin.H{"access_token": token}, http.StatusOK)
	jsonResponse := utils.MakeResponse(httpResponse, "Success")
	c.JSON(httpResponse.Status, jsonResponse)
}

func NewUserController(usecase IUserUseCase) *userController {
	return &userController{
		UseCases: usecase,
	}
}
