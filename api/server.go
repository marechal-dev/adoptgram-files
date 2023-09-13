package api

import (
	"github.com/gin-gonic/gin"
)

const (
	ServerDefaultAddress = "0.0.0.0"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{}

	server.setupRouter()

	return server
}

func (s *Server) setupRouter() {
	router := gin.Default()

	// Routes here

	s.router = router
}

func (s *Server) SpinUp(addr string) {
	s.router.Run(addr)
}
