package handler

import (
	"golang-chat-backend/json/response"

	"github.com/gin-gonic/gin"
)

func NoRouteHandler(ctx *gin.Context) {
	ctx.JSON(404, response.BaseResponse{
		Message: "",
		Data:    "",
		Error:   "404 Endpoint not found",
	})
}

func NoMethodAllowed(ctx *gin.Context) {
	ctx.JSON(400, response.BaseResponse{
		Message: "",
		Data:    "",
		Error:   "No Method Allowed",
	})
}

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, response.BaseResponse{
		Message: "",
		Data:    "",
		Error:   "Chat Message API is Up",
	})
}
