package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{}

	server.setupRouter()

	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// Routes here

	server.router = router
}
