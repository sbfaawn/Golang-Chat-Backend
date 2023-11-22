package router

import (
	"golang-chat-backend/authentication"
	"golang-chat-backend/handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func (router *Router) InitalizeRouter() {
	r := router.engine

	r.Use(gin.Recovery())
	r.NoRoute(handler.NoRouteHandler)
	r.NoMethod(handler.NoMethodAllowed)

	group := r.Group("", authentication.BasicAuth).Group("/api")
	account := group.Group("/account")

	// health check
	group.GET("/health", handler.HealthCheck)

	// account
	account.POST("/register", handler.RegistrationHandler)
	account.POST("/login", handler.LoginHandler)
	account.POST("/logout", handler.LogoutHandler)
	account.POST("/refresh", handler.RefreshTokenHandler)
}

func (router *Router) Start(port string) {
	router.engine.Run(":" + port)
}

func NewRouter() *Router {
	return &Router{
		engine: gin.Default(),
	}
}
