package server

import (
	"golang-chat-backend/api/authentication"
	"golang-chat-backend/api/handler"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func (s *Server) InitalizeServer() {
	server := s.engine

	server.Use(gin.Recovery())
	server.NoRoute(handler.NoRouteHandler)
	server.NoMethod(handler.NoMethodAllowed)

	group := server.Group("", authentication.BasicAuth).Group("/api")
	account := group.Group("/account")

	// health check
	group.GET("/health", handler.HealthCheck)

	// account
	account.POST("/register", handler.RegistrationHandler)
	account.POST("/login", handler.LoginHandler)
	account.POST("/logout", handler.LogoutHandler)
	account.POST("/refresh", handler.RefreshTokenHandler)
}

func (s *Server) Start(port string) {
	s.engine.Run(":" + port)
}

func NewServer() *Server {
	return &Server{
		engine: gin.Default(),
	}
}
