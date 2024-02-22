package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/yangoneseok/voyager/db/sqlc"
)

// Server serves HTTP requests for app service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}

	router := gin.Default()
	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUser)
	router.GET("/users", server.listUser)
	router.PATCH("/users", server.updateUser)
	router.DELETE("/users", server.deleteUser)

	server.router = router

	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
