package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func CreateHttpServer() *Server {
	return &Server{
		router: gin.Default(),
	}
}
func (s *Server) RegisterHandler(path, method string, handler gin.HandlerFunc) {
	switch method {
	case "GET":
		s.router.GET(path, handler)
	case "POST":
		s.router.POST(path, handler)
	}
}

func (s *Server) StartServer(port string) error {
	err := s.router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}
	return nil
}
