package authentication

import (
	"golang-chat-backend/models/output"

	"github.com/gin-gonic/gin"
)

func BasicAuth(ctx *gin.Context) {
	username, password, isOk := ctx.Request.BasicAuth()

	if !(isOk && username == "chatuser" && password == "HolE34@HJ") {
		ctx.JSON(401, output.BaseResponse{
			Error:   "",
			Data:    "",
			Message: "unauthorized user, username and password need to access this resource",
		})
		ctx.Abort()
		ctx.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		return
	}
}
