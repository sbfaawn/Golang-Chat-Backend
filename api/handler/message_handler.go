package handler

import (
	"errors"
	"golang-chat-backend/models"
	"golang-chat-backend/models/input"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *HttpHandler) SendMessages(ctx *gin.Context) {
	input := input.MessageInput{}
	if err := ctx.BindJSON(&input); err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	if err := h.JsonValidator.Validate(&input); err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	message := models.Message{
		Sender:    input.Sender,
		Receiver:  input.Receiver,
		Message:   input.Message,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := h.messageService.SendMessage(ctx, &message)

	if err == nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	generateResponse(ctx, 200, "Message has beeen sent", nil)
}

func (h *HttpHandler) GetConversation(ctx *gin.Context) {
	sender := ctx.Query("sender")
	receiver := ctx.Query("receiver")

	if sender == "" || receiver == "" {
		generateResponse(ctx, 400, "", errors.New("sender or receiver is not specified on request"))
	}

	messages, err := h.messageService.GetConversation(ctx, sender, receiver)

	if err == nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	generateResponse(ctx, 200, messages, nil)
}

func (h *HttpHandler) DeleteMessage(ctx *gin.Context) {
	messageId := ctx.Query("id")
	if messageId == "" {
		generateResponse(ctx, 400, "", errors.New("message id is not specified"))
		return
	}

	id, err := strconv.Atoi(messageId)
	if err != nil {
		generateResponse(ctx, 400, "", err)
		return
	}

	h.messageService.DeleteMessage(ctx, id)

	generateResponse(ctx, 200, "Message has beeen deleted", nil)
}
