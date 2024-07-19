package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	config "github.com/resistance22/university_project/Config"
	controller "github.com/resistance22/university_project/Controller"
	repository "github.com/resistance22/university_project/Repository"
	token "github.com/resistance22/university_project/Token"
	usecase "github.com/resistance22/university_project/UseCase"
	db "github.com/resistance22/university_project/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(config *config.Config, store *db.Store) *Server {
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
			tokenMaker, err := token.NewPasteoTokenMaker([]byte(config.TokenKey))
			if err != nil {
				log.Fatal(err.Error())
			}
			repository := repository.NewUserRepository(store)
			userUseCase := usecase.NewUserUseCase(repository, tokenMaker)
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
