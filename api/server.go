package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/resistance22/university_project/db/sqlc"
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

	server := &Server{
		store:  store,
		router: router,
	}

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
