package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/resistance22/university_project/Validator"
	usecase "github.com/resistance22/university_project/usecase"
	"github.com/resistance22/university_project/utils"
)

type userController struct {
	UseCases *usecase.UserUseCase
}

func (controller *userController) Register(c *gin.Context) {
	var json validator.RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		httpError := utils.NewHttpError(err.Error(), http.StatusBadRequest, "Invalid Body")
		response := utils.MakeError(httpError)
		c.JSON(httpError.Status, response)
		return
	}
}
