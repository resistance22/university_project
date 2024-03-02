package server

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router    *gin.Engine
	DbContext context.Context
}

func NewServer(dbContext context.Context) *Server {
	server := &Server{}
	router := gin.Default()
	router.GET("/ping", server.ping)
	server.router = router
	return server
}

func (server *Server) Start(address string) {
	server.router.Run(address)
}

// func errorResponse(err error) gin.H {
// 	return gin.H{"error": err.Error()}
// }
