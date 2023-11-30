package server

import (
	"golang-chat-backend/api/authentication"
	"golang-chat-backend/api/handler"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine      *gin.Engine
	httpHandler *handler.HttpHandler
}

func NewServer() *Server {
	return &Server{
		engine: gin.Default(),
	}
}

func (s *Server) InitalizeServer() {
	server := s.engine

	server.Use(gin.Recovery())
	server.NoRoute(s.httpHandler.NoRouteHandler)
	server.NoMethod(s.httpHandler.NoMethodAllowed)

	group := server.Group("", authentication.CheckBasicAuth).Group("/api")
	account := group.Group("/account")

	// health check
	group.GET("/health", s.httpHandler.HealthCheck)

	// account
	account.POST("/register", s.httpHandler.RegistrationHandler)
	account.POST("/login", s.httpHandler.LoginHandler)
	account.POST("/logout", s.httpHandler.LogoutHandler)
	account.POST("/refresh", s.httpHandler.RefreshTokenHandler)
}

func (s *Server) Start(port string) {
	s.engine.Run(":" + port)
}
