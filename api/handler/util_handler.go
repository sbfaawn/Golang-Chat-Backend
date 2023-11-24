package handler

import (
	"golang-chat-backend/models/output"

	"github.com/gin-gonic/gin"
)

func generateResponse(ctx *gin.Context, statusCode int, data any, err error) {
	var message string

	message = "Success"
	if statusCode != 200 {
		message = "Failed"
	}

	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}

	ctx.JSON(statusCode, output.BaseResponse{
		Message: message,
		Data:    data,
		Error:   errorMessage,
	})
}
