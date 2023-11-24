package storage

import (
	"golang-chat-backend/models"

	"github.com/gin-gonic/gin"
)

type AccountStorageInterface interface {
	SaveAccount(ctx *gin.Context, account models.Account) error
	GetAccountByUsername(ctx *gin.Context, username string) (models.Account, error)
	GetAccountByEmail(ctx *gin.Context, email string) (models.Account, error)
	UpdatePasswordByUsername(ctx *gin.Context, username string, newPassword string) error
	UpdateVerifiedByEmail(ctx *gin.Context, email string) error
}
