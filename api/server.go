package api

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/yangoneseok/voyager/docs"

	db "github.com/yangoneseok/voyager/db/sqlc"
	"github.com/yangoneseok/voyager/db/token"
	"github.com/yangoneseok/voyager/util"
)

// Server serves HTTP requests for app service
type Server struct {
	config     util.Config
	logger     *log.Logger
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
	logger := log.New(os.Stdout, "voyager ", log.LstdFlags)
	server := &Server{
		config:     config,
		logger:     logger,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	if server.config.SwaggerEnabled {
		router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	server.setupAuth(router)
	server.setupUsers(router)

	server.setupTokens(router)

	server.router = router
}

func (server *Server) setupAuth(router *gin.Engine) {
	router.POST("/auth/login", server.loginUser)
	router.POST("/auth/signup", server.signUpUser)

	authRoutes := router.Group("/").
		Use(server.authMiddleware())

	authRoutes.GET("/me", server.getMe)
}

func (server *Server) setupUsers(router *gin.Engine) {
	userGroup := router.Group("/users")

	userGroup.POST("", server.createUser)
	userGroup.GET("/:id", server.getUser)
	userGroup.GET("", server.listUser)
	userGroup.PATCH("", server.updateUser)
	userGroup.DELETE("", server.deleteUser)
}

func (server *Server) setupTokens(router *gin.Engine) {
	router.POST("/tokens/reissue", server.renewAccessToken)
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
