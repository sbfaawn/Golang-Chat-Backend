package service

import (
	"golang-chat-backend/models"
	"golang-chat-backend/storage"
	"golang-chat-backend/util"

	"github.com/gin-gonic/gin"
)

type AccountServiceInterface interface {
	SaveAccount(ctx *gin.Context, account *models.Account) error
	ChangePassword(ctx *gin.Context, account *models.Account) error
	AccountVerification(ctx *gin.Context, account *models.Account) (models.Account, error)
	Login(ctx *gin.Context, account *models.Account) error
}

type accountService struct {
	accountStorage storage.AccountStorageInterface
	encryptor      *util.PasswordEncryptor
}

func NewAccountService(accountStorage storage.AccountStorageInterface, encryptor *util.PasswordEncryptor) *accountService {
	return &accountService{
		accountStorage: accountStorage,
		encryptor:      encryptor,
	}
}

func (service *accountService) SaveAccount(ctx *gin.Context, account *models.Account) error {
	return nil
}

func (service *accountService) ChangePassword(ctx *gin.Context, account *models.Account) error {
	return nil
}

func (service *accountService) AccountVerification(ctx *gin.Context, account *models.Account) (models.Account, error) {
	return models.Account{}, nil
}

func (service *accountService) Login(ctx *gin.Context, account *models.Account) error {
	return nil
}
