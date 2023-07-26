package server

import (
	"api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5555",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes((s.server))

	log.Print("Server is running on port 5555")
	log.Fatal(router.Run(":" + s.port))
}
