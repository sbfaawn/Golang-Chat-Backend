package storage

import (
	"golang-chat-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SessionStorageInterface interface {
	SaveSession(ctx *gin.Context, session *models.Session) error
	GetSessionById(ctx *gin.Context, sessionId string) (models.Session, error)
}

type SessionStorage struct {
	DB *gorm.DB
}

func NewSessionStorage(DB *gorm.DB) *SessionStorage {
	return &SessionStorage{
		DB: DB,
	}
}

func (storage *SessionStorage) SaveSession(ctx *gin.Context, session *models.Session) error {
	var err error
	db := storage.DB

	tx := db.Begin()
	err = tx.Create(&session).Error

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

func (storage *SessionStorage) GetSessionById(ctx *gin.Context, sessionId string) (models.Session, error) {
	db := storage.DB
	var err error
	var session models.Session

	tx := db.Begin()
	err = tx.First(&session, "id = ?", sessionId).Error

	if err != nil {
		tx.Rollback()
		return session, err
	}

	err = tx.Commit().WithContext(ctx).Error

	if err != nil {
		tx.Rollback()
		return session, err
	}

	return session, nil
}
