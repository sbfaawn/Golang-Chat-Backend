package storage

import (
	"errors"
	"golang-chat-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SessionStorageInterface interface {
	SaveSession(ctx *gin.Context, session *models.Session) error
	GetSessionById(ctx *gin.Context, sessionId string) (models.Session, error)
	DeleteSession(ctx *gin.Context, sessionId string) error
	UpdateSessionExpiration(ctx *gin.Context, session *models.Session) error
}

type SessionStorage struct {
	db *gorm.DB
}

func NewSessionStorage(DB *gorm.DB) *SessionStorage {
	return &SessionStorage{
		db: DB,
	}
}

func (storage *SessionStorage) SaveSession(ctx *gin.Context, session *models.Session) error {
	var err error
	db := storage.db

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
	db := storage.db
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

func (storage *SessionStorage) DeleteSession(ctx *gin.Context, sessionId string) error {
	var err error
	db := storage.db
	session := models.Session{
		Id: sessionId,
	}

	tx := db.Begin()
	delete := tx.Model(&session).Delete(&session)

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
		return errors.New("record with sessionID is not found")
	}

	return nil
}

func (storage *SessionStorage) UpdateSessionExpiration(ctx *gin.Context, session *models.Session) error {
	db := storage.db
	var err error

	tx := db.Begin()
	update := tx.Model(&models.Account{}).Where("id = ? AND username = ?", session.Id, session.Username).Updates(map[string]any{
		"expired_at": session.ExpiredAt,
	})

	if err != nil {
		tx.Rollback()
		return err
	}

	result := update.Commit().WithContext(ctx)

	if result.Error != nil {
		tx.Rollback()
		return err
	}

	if update.RowsAffected == result.RowsAffected {
		tx.Rollback()
		return errors.New("record with sessionId and username is not found")
	}

	return nil
}
