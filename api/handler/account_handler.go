package handler

import (
	"golang-chat-backend/models"

	"github.com/gin-gonic/gin"
)

const (
	tokenLifetime int    = 5
	jwtTokenKey   string = "jwt-token"
)

func RegistrationHandler(ctx *gin.Context) {
	account := models.Account{}
	if err := ctx.BindJSON(&account); err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

}

func LoginHandler(ctx *gin.Context) {

}

func LogoutHandler(ctx *gin.Context) {
	ctx.SetCookie(jwtTokenKey, "", -1, "/", "localhost", false, true)
	generateResponse(ctx, 200, "", nil)
}

func RefreshTokenHandler(ctx *gin.Context) {

}
