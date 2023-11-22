package handler

import "github.com/gin-gonic/gin"

const (
	tokenLifetime int    = 5
	jwtTokenKey   string = "jwt-token"
)

func RegistrationHandler(ctx *gin.Context) {

}

func LoginHandler(ctx *gin.Context) {

}

func LogoutHandler(ctx *gin.Context) {
	ctx.SetCookie(jwtTokenKey, "", -1, "/", "localhost", false, true)
	generateResponse(ctx, 200, "", nil)
}

func RefreshTokenHandler(ctx *gin.Context) {

}
