package storage

import (
	"errors"
	"golang-chat-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountStorageInterface interface {
	SaveAccount(ctx *gin.Context, account models.Account) error
	GetAccountByUsername(ctx *gin.Context, username string) (models.Account, error)
	GetAccountByEmail(ctx *gin.Context, email string) (models.Account, error)
	UpdatePasswordByUsername(ctx *gin.Context, username string, newPassword string) error
	UpdateVerifiedByEmail(ctx *gin.Context, email string) error
}

type accountStorage struct {
	DB *gorm.DB
}

func NewAccountStorage(DB *gorm.DB) *accountStorage {
	return &accountStorage{
		DB: DB,
	}
}

func (storage *accountStorage) SaveAccount(ctx *gin.Context, account models.Account) error {
	var err error
	db := storage.DB

	tx := db.Begin()
	err = tx.Create(&account).Error

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

func (storage *accountStorage) GetAccountByUsername(ctx *gin.Context, username string) (models.Account, error) {
	db := storage.DB
	var err error
	var account models.Account

	tx := db.Begin()
	err = tx.First(&account, "username = ?", username).Error

	if err != nil {
		tx.Rollback()
		return account, err
	}

	err = tx.Commit().WithContext(ctx).Error

	if err != nil {
		tx.Rollback()
		return account, err
	}

	return account, nil
}

func (storage *accountStorage) UpdatePasswordByUsername(ctx *gin.Context, username string, newPassword string) error {
	db := storage.DB
	var err error

	tx := db.Begin()
	update := tx.Model(&models.Account{}).Where("username = ? AND deleted_at IS null", username).Updates(map[string]any{
		"password": newPassword,
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
		return errors.New("record with username " + username + " is not found")
	}

	return nil
}

func (storage *accountStorage) GetAccountByEmail(ctx *gin.Context, email string) (models.Account, error) {
	db := storage.DB
	var err error
	var account models.Account

	tx := db.Begin()
	err = tx.First(&account, "email = ?", email).Error

	if err != nil {
		tx.Rollback()
		return account, err
	}

	err = tx.Commit().WithContext(ctx).Error

	if err != nil {
		tx.Rollback()
		return account, err
	}

	return account, nil
}

func (storage *accountStorage) UpdateVerifiedByEmail(ctx *gin.Context, email string) error {
	db := storage.DB
	var err error

	tx := db.Begin()
	update := tx.Model(&models.Account{}).Where("email = ? AND deleted_at IS null", email).Updates(map[string]any{
		"verified": true,
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
		return errors.New("record with email " + email + " is not found")
	}

	return nil
}
