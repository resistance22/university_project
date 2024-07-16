package utils

import "github.com/gin-gonic/gin"

func MakeResponse(arg *HttpResponse, message string) *gin.H {
	return &gin.H{
		"success": true,
		"result":  arg,
		"message": message,
	}
}

func MakeError(err *HttpError) *gin.H {
	return &gin.H{
		"success": false,
		"message": err.Message,
		"result":  err.Description,
	}
}
