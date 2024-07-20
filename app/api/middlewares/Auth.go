package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	token "github.com/resistance22/university_project/Token"
	utils "github.com/resistance22/university_project/Utils"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "user_payload"
)

func AuthMiddleware(maker token.TokenMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			httpError := utils.NewHttpError("unauthorized", http.StatusUnauthorized, "No Auth Token Provided!")
			response := utils.MakeError(httpError)
			ctx.AbortWithStatusJSON(httpError.Status, response)
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			httpError := utils.NewHttpError("unauthorized", http.StatusUnauthorized, "Invalid Auth Token Format!")
			response := utils.MakeError(httpError)
			ctx.AbortWithStatusJSON(httpError.Status, response)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			httpError := utils.NewHttpError("unauthorized", http.StatusUnauthorized, "Invalid Auth Type!")
			response := utils.MakeError(httpError)
			ctx.AbortWithStatusJSON(httpError.Status, response)
			return
		}

		accessToken := fields[1]
		payload, err := maker.VerifyToken(accessToken)

		if err != nil {
			httpError := utils.NewHttpError("unauthorized", http.StatusUnauthorized, err.Error())
			response := utils.MakeError(httpError)
			ctx.AbortWithStatusJSON(httpError.Status, response)
			return
		}

		if payload.Valid() != nil {
			httpError := utils.NewHttpError("unauthorized", http.StatusUnauthorized, "Token Expired!")
			response := utils.MakeError(httpError)
			ctx.AbortWithStatusJSON(httpError.Status, response)
			return
		}

		ctx.Set(authorizationPayloadKey, payload)

	}
}
