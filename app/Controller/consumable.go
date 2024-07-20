package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/resistance22/university_project/UseCase"
	utils "github.com/resistance22/university_project/Utils"
	validator "github.com/resistance22/university_project/Validator"
)

type consumableController struct {
	UseCases usecase.IConsumableUseCase
}

func (controller *consumableController) Create(c *gin.Context) {
	var body validator.CreateConsumableBody
	if err := c.ShouldBindJSON(&body); err != nil {
		httpError := utils.NewHttpError(err.Error(), http.StatusBadRequest, "Invalid Body")
		response := utils.MakeError(httpError)
		c.JSON(httpError.Status, response)
		return
	}

	res, err := controller.UseCases.Create(c, &body)

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

	httpResponse := utils.NewHttpResponse(response, http.StatusCreated)
	jsonResponse := utils.MakeResponse(httpResponse, "New User Created")
	c.JSON(httpResponse.Status, jsonResponse)
}

func NewConsumableController(usecase usecase.IConsumableUseCase) *consumableController {
	return &consumableController{
		UseCases: usecase,
	}
}
