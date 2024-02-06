package storage

import (
	"errors"
	"golang-chat-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MessageStorageInterface interface {
	SaveMessage(ctx *gin.Context, message *models.Message) error
	GetMessagesByConversation(ctx *gin.Context, usernameOne string, usernameTwo string) ([]models.Message, error)
	DeleteMessageById(ctx *gin.Context, messageId int) error
}

type MessageStorage struct {
	db *gorm.DB
}

func NewMessageStorage(DB *gorm.DB) *MessageStorage {
	return &MessageStorage{
		db: DB,
	}
}

func (storage *MessageStorage) SaveMessage(ctx *gin.Context, message *models.Message) error {
	var err error
	db := storage.db

	tx := db.Begin()
	err = tx.Create(&message).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().WithContext(ctx).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (storage *MessageStorage) GetMessagesByConversation(ctx *gin.Context, usernameOne string, usernameTwo string) ([]models.Message, error) {
	db := storage.db
	var err error
	var messages []models.Message

	tx := db.Begin()
	err = tx.Select("message_id", "sender", "receiver", "message", "created_at").
		Where("(sender = ? AND receiver = ?) OR (sender = ? AND receiver = ?)", usernameOne, usernameTwo, usernameTwo, usernameOne).
		Find(&messages).
		Error

	if err != nil {
		tx.Rollback()
		return messages, err
	}

	err = tx.Commit().WithContext(ctx).Error

	if err != nil {
		tx.Rollback()
		return messages, err
	}

	return messages, nil
}

func (storage *MessageStorage) DeleteMessageById(ctx *gin.Context, messageId int) error {
	var err error
	db := storage.db
	message := models.Message{
		Id: messageId,
	}

	tx := db.Begin()
	delete := tx.Model(&message).Delete(&message)

	if delete.Error != nil {
		tx.Rollback()
		return err
	}

	result := delete.Commit().WithContext(ctx)

	if result.Error != nil {
		return err
	}

	if delete.RowsAffected == result.RowsAffected {
		tx.Rollback()
		return errors.New("record with messageId is not found")
	}

	return nil
}
