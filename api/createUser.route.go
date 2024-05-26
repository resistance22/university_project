package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/resistance22/micorsales/db/sqlc"
)

type RegisterUserBodyValidator struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req RegisterUserBodyValidator

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, server.errorResponse(err))
		return
	}

	newUserData := db.CreateUserParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Role:      "admin",
	}

	newUser, err := server.queries.CreateUser(server.DbContext, newUserData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, server.errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, newUser)
}
