package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/resistance22/university_project/utils"
)

func translateToHttpError(err any) *utils.HttpError {
	switch v := err.(type) {
	case utils.HttpError:
		return &v
	default:
		return &utils.HttpError{
			Message: "Something Went Wrong!",
			Status:  http.StatusInternalServerError,
		}
	}
}

func ErrorHandler(c *gin.Context, err any) {
	httpError := translateToHttpError(err)
	err = utils.MakeError(httpError)
	c.JSON(httpError.Status, err)
}
