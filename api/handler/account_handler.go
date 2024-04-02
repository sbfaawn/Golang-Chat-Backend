package handler

import (
	"errors"
	"fmt"
	"golang-chat-backend/models"
	"golang-chat-backend/models/input"

	"github.com/gin-gonic/gin"
)

func (h *HttpHandler) RegistrationHandler(ctx *gin.Context) {
	input := input.AccountInput{}
	if err := ctx.BindJSON(&input); err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	if input.Email == "" {
		generateResponse(ctx, 400, "Email is needed in registration process", nil)
		return
	}

	if err := h.JsonValidator.Validate(&input); err != nil {
		generateResponse(ctx, 400, "Input is not Valid", err)
		return
	}

	account := models.Account{
		Username:   input.Username,
		Password:   input.Password,
		Email:      input.Email,
		IsVerified: false,
	}

	err := h.accountservice.SaveAccount(ctx, &account)

	if err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	generateResponse(ctx, 200, "Account is register succesfully", nil)
}

func (h *HttpHandler) LoginHandler(ctx *gin.Context) {
	input := input.AccountInput{}
	if err := ctx.BindJSON(&input); err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	if err := h.JsonValidator.Validate(&input); err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	account := models.Account{
		Username: input.Username,
		Password: input.Password,
	}

	err := h.accountservice.Login(ctx, &account)

	// user
	if err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	session, err2 := h.sessionService.CreateSession(ctx, account.Username)

	if err2 != nil {
		generateResponse(ctx, 400, "", err2)
		return
	}

	fmt.Println(session.TTL)
	fmt.Println(session.TTL.Seconds())
	ctx.SetCookie("SessionID", session.Id, int(session.TTL.Seconds()), "/", "localhost", false, true)
	generateResponse(ctx, 200, "Login Successfully", nil)
}

func (h *HttpHandler) LogoutHandler(ctx *gin.Context) {
	sessionId, err := ctx.Cookie("SessionID")

	if err != nil {
		generateResponse(ctx, 404, "", err)
		return
	}

	err = h.sessionService.RemoveSession(ctx, sessionId)

	if err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	ctx.SetCookie("SessionID", "", -1, "/", "localhost", false, true)
	generateResponse(ctx, 200, "", nil)
}

func (h *HttpHandler) RefreshTokenHandler(ctx *gin.Context) {
	sessionId, err := ctx.Cookie("SessionID")

	if err != nil {
		generateResponse(ctx, 404, "", err)
		return
	}

	session, err := h.sessionService.UpdateSessionExpiration(ctx, sessionId)

	if err != nil {
		generateResponse(ctx, 404, "", err)
		return
	}

	ctx.SetCookie("SessionID", session.Id, int(session.TTL.Seconds()), "/", "localhost", false, true)
	generateResponse(ctx, 200, "Token has been refreshed", nil)
}

func (h *HttpHandler) CheckSession(ctx *gin.Context) {
	var err error
	sessionId, err := ctx.Cookie("SessionID")

	if err != nil {
		generateResponse(ctx, 404, "", err)
		ctx.AbortWithError(404, errors.New("session is not included on request"))
		return
	}

	_, err = h.sessionService.CheckSession(ctx, sessionId)
	if err != nil {
		generateResponse(ctx, 404, "session id is not found", err)
		ctx.AbortWithError(404, errors.New("session is not found, please input valid session ID or login"))
		return
	}
}
