package handler

import (
	"golang-chat-backend/service"
	"golang-chat-backend/util"
)

type HttpHandler struct {
	*util.JsonValidator
	accountservice service.AccountServiceInterface
	sessionService service.SessionServiceInterface
	messageService service.MessageServiceInterface
}

func NewHttpHandler(
	accountService service.AccountServiceInterface,
	sessionService service.SessionServiceInterface,
	messageService service.MessageServiceInterface,
) *HttpHandler {
	return &HttpHandler{
		JsonValidator:  util.NewJsonValidator(),
		accountservice: accountService,
		sessionService: sessionService,
		messageService: messageService,
	}
}
