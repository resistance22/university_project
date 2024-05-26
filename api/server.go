package server

import (
	"context"

	"github.com/gin-gonic/gin"
	db "github.com/resistance22/micorsales/db/sqlc"
)

type Server struct {
	router    *gin.Engine
	DbContext context.Context
	queries   *db.Queries
}

func NewServer(dbContext context.Context, queries *db.Queries) *Server {
	server := &Server{}
	router := gin.Default()
	router.GET("/ping", server.ping)
	router.POST("/register", server.CreateUser)
	server.router = router
	server.DbContext = dbContext
	server.queries = queries
	return server
}

func (server *Server) Start(address string) {
	server.router.Run(address)
}

func (s *Server) errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
