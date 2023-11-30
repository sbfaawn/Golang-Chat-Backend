package handler

import (
	"golang-chat-backend/service"
	"golang-chat-backend/util"
)

type HttpHandler struct {
	*util.JsonValidator
	accountservice service.AccountServiceInterface
}

func NewHttpHandler(accountService service.AccountServiceInterface) *HttpHandler {
	return &HttpHandler{
		JsonValidator:  util.NewJsonValidator(),
		accountservice: accountService,
	}
}
