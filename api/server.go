package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	repository "github.com/resistance22/university_project/Repository"
	"github.com/resistance22/university_project/controller"
	db "github.com/resistance22/university_project/db/sqlc"
	"github.com/resistance22/university_project/usecase"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result": "pong",
		})
	})

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("auth")
		{
			repository := repository.NewUserRepository(store)
			userUseCase := usecase.NewUserUseCase(repository)
			controller := controller.NewUserController(userUseCase)
			auth.POST("/register", controller.Register)
		}
	}

	server := &Server{
		store:  store,
		router: router,
	}

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
