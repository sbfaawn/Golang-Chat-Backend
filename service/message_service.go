package service

import (
	"golang-chat-backend/models"
	"golang-chat-backend/storage"

	"github.com/gin-gonic/gin"
)

type MessageServiceInterface interface {
	SendMessage(ctx *gin.Context, message *models.Message) error
	GetConversation(ctx *gin.Context, usernameOne string, usernameTwo string) ([]models.Message, error)
	DeleteMessage(ctx *gin.Context, messageId int) error
}

type MessageService struct {
	messageStorage storage.MessageStorageInterface
}

func NewMessageService(messageStorage storage.MessageStorageInterface) *MessageService {
	return &MessageService{
		messageStorage: messageStorage,
	}
}

func (service *MessageService) SendMessage(ctx *gin.Context, message *models.Message) error {
	err := service.messageStorage.SaveMessage(ctx, message)
	if err != nil {
		return err
	}

	return nil
}

func (service *MessageService) GetConversation(ctx *gin.Context, usernameOne string, usernameTwo string) ([]models.Message, error) {
	messages, err := service.messageStorage.GetMessagesByConversation(ctx, usernameOne, usernameTwo)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (service *MessageService) DeleteMessage(ctx *gin.Context, messageId int) error {
	err := service.messageStorage.DeleteMessageById(ctx, messageId)
	if err != nil {
		return err
	}

	return nil
}
