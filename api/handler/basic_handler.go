package handler

import (
	"golang-chat-backend/models/output"

	"github.com/gin-gonic/gin"
)

func (h *HttpHandler) NoRouteHandler(ctx *gin.Context) {
	ctx.JSON(404, output.BaseResponse{
		Message: "",
		Data:    "",
		Error:   "404 Endpoint not found",
	})
}

func (h *HttpHandler) NoMethodAllowed(ctx *gin.Context) {
	ctx.JSON(400, output.BaseResponse{
		Message: "",
		Data:    "",
		Error:   "No Method Allowed",
	})
}

func (h *HttpHandler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, output.BaseResponse{
		Message: "",
		Data:    "",
		Error:   "Chat Message API is Up",
	})
}
