package service

import (
	"errors"
	"golang-chat-backend/models"
	"golang-chat-backend/storage"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	SESSION_TTL = 60 * time.Minute
)

type SessionServiceInterface interface {
	CreateSession(ctx *gin.Context, username string) (models.Session, error)
	CheckSession(ctx *gin.Context, sessionId string) (string, error)
	RemoveSession(ctx *gin.Context, sessionId string) error
	UpdateSessionExpiration(ctx *gin.Context, sessionId string) (models.Session, error)
}

type SessionService struct {
	sessionStorage storage.SessionStorageInterface
}

func NewSessionService(sessionStorage storage.SessionStorageInterface) *SessionService {
	return &SessionService{
		sessionStorage: sessionStorage,
	}
}

func (s *SessionService) CreateSession(ctx *gin.Context, username string) (models.Session, error) {
	sessionToken := uuid.NewString()

	session := models.Session{
		Id:       sessionToken,
		Username: username,
		TTL:      SESSION_TTL,
	}

	err := s.sessionStorage.SaveSession(ctx, &session)

	if err != nil {
		return models.Session{}, err
	}

	return session, nil
}

func (s *SessionService) CheckSession(ctx *gin.Context, sessionId string) (string, error) {
	username, err := s.sessionStorage.GetSessionById(ctx, sessionId)

	if err != nil {
		return "", errors.New("Session with SessionId " + sessionId + " is not found")
	}

	return username, nil
}

func (s *SessionService) RemoveSession(ctx *gin.Context, sessionId string) error {
	err := s.sessionStorage.DeleteSession(ctx, sessionId)

	if err != nil {
		return err
	}

	return nil
}

func (s *SessionService) UpdateSessionExpiration(ctx *gin.Context, sessionId string) (models.Session, error) {
	username, err := s.sessionStorage.GetSessionById(ctx, sessionId)

	if err != nil {
		return models.Session{}, errors.New("Session with SessionId " + sessionId + " is not found")
	}

	session := models.Session{
		Username: username,
		TTL:      SESSION_TTL,
	}

	err = s.sessionStorage.UpdateSessionExpiration(ctx, &session)

	if err != nil {
		return models.Session{}, err
	}

	return session, nil
}
