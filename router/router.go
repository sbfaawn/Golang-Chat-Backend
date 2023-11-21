package router

import (
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

	group := r.Group("/api")

	group.GET("/health", handler.HealthCheck)
}

func (router *Router) Start(port string) {
	router.engine.Run(":" + port)
}

func NewRouter() *Router {
	return &Router{
		engine: gin.Default(),
	}
}
