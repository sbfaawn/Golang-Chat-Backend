package handler

import (
	"golang-chat-backend/service"
	"golang-chat-backend/util"
)

type HttpHandler struct {
	*util.JsonValidator
	accountservice service.AccountServiceInterface
	sessionService service.SessionServiceInterface
}

func NewHttpHandler(
	accountService service.AccountServiceInterface,
	sessionService service.SessionServiceInterface,
) *HttpHandler {
	return &HttpHandler{
		JsonValidator:  util.NewJsonValidator(),
		accountservice: accountService,
		sessionService: sessionService,
	}
}
