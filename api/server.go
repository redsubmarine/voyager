package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/yangoneseok/voyager/db/sqlc"
	"github.com/yangoneseok/voyager/db/token"
	"github.com/yangoneseok/voyager/util"
)

// Server serves HTTP requests for app service
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	authRoutes := router.Group("/").
		Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/me", server.getMe)

	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUser)
	router.GET("/users", server.listUser)
	router.PATCH("/users", server.updateUser)
	router.DELETE("/users", server.deleteUser)

	router.POST("/tokens/renew", server.renewAccessToken)

	// auth
	router.POST("/auth/login", server.loginUser)
	server.router = router
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
