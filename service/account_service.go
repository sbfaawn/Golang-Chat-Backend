package service

import (
	"errors"
	"golang-chat-backend/models"
	"golang-chat-backend/storage"
	"golang-chat-backend/util"

	"github.com/gin-gonic/gin"
)

type AccountServiceInterface interface {
	SaveAccount(ctx *gin.Context, account *models.Account) error
	ChangePassword(ctx *gin.Context, account *models.Account) error
	AccountVerification(ctx *gin.Context, account *models.Account) error
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

func (s *accountService) SaveAccount(ctx *gin.Context, account *models.Account) error {
	_, err := s.accountStorage.GetAccountByUsername(ctx, account.Username)

	if err == nil {
		return errors.New("Account with username " + account.Username + " is already exist")
	}

	s.encryptor.Password = account.Password
	passEncrypt, err := s.encryptor.Encrypt()

	if err != nil {
		return errors.New("Error is occured when encrypt password")
	}

	account.Password = passEncrypt
	err = s.accountStorage.SaveAccount(ctx, account)

	if err != nil {
		return errors.New("Error when trying create an account")
	}

	return nil
}

func (s *accountService) ChangePassword(ctx *gin.Context, account *models.Account) error {
	err := s.accountStorage.UpdatePasswordByUsername(ctx, account.Username, account.Password)
	if err != nil {
		return err
	}

	return nil
}

func (s *accountService) AccountVerification(ctx *gin.Context, account *models.Account) error {
	err := s.accountStorage.UpdateVerifiedByEmail(ctx, account.Email)
	if err != nil {
		return err
	}

	return nil
}

func (s *accountService) Login(ctx *gin.Context, account *models.Account) error {
	_, err := s.accountStorage.GetAccountByUsername(ctx, account.Username)

	if err != nil {
		return errors.New("Account with username " + account.Username + " is not exist")
	}

	return nil
}
