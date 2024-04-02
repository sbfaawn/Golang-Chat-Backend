package storage

import (
	"golang-chat-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type SessionStorageInterface interface {
	SaveSession(ctx *gin.Context, session *models.Session) error
	GetSessionById(ctx *gin.Context, sessionId string) (string, error)
	DeleteSession(ctx *gin.Context, sessionId string) error
	UpdateSessionExpiration(ctx *gin.Context, session *models.Session) error
}

type SessionStorage struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewSessionStorage(DB *gorm.DB, redisClient *redis.Client) *SessionStorage {
	return &SessionStorage{
		db:          DB,
		redisClient: redisClient,
	}
}

func (storage *SessionStorage) SaveSession(ctx *gin.Context, session *models.Session) error {
	var err error
	redisClient := storage.redisClient

	setCmd := redisClient.Set(ctx, session.Id, session.Username, session.TTL)
	if err = setCmd.Err(); err != nil {
		return err
	}

	return nil
}

func (storage *SessionStorage) GetSessionById(ctx *gin.Context, sessionId string) (string, error) {
	var err error
	redisClient := storage.redisClient

	getCmd := redisClient.Get(ctx, sessionId)
	if err = getCmd.Err(); err != nil {
		return "", err
	}

	res, err := getCmd.Result()
	if err != nil {
		return "", err
	}

	return res, nil
}

func (storage *SessionStorage) DeleteSession(ctx *gin.Context, sessionId string) error {
	var err error
	redisClient := storage.redisClient

	delCmd := redisClient.Del(ctx, sessionId)
	if err = delCmd.Err(); err != nil {
		return err
	}

	return nil
}

func (storage *SessionStorage) UpdateSessionExpiration(ctx *gin.Context, session *models.Session) error {
	var err error
	redisClient := storage.redisClient

	setCmd := redisClient.Expire(ctx, session.Id, session.TTL)
	if err = setCmd.Err(); err != nil {
		return err
	}

	return nil
}
