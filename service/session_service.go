package service

import (
	"database/sql"
	"errors"
	"golang-chat-backend/models"
	"golang-chat-backend/storage"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	SESSION_TTL = 120
)

type SessionServiceInterface interface {
	CreateSession(ctx *gin.Context, username string) (models.Session, error)
	CheckSession(ctx *gin.Context, sessionId string) error
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
	expiredAt := time.Now().Add(time.Minute * SESSION_TTL)

	session := models.Session{
		Id:       sessionToken,
		Username: username,
		ExpiredAt: sql.NullTime{
			Time:  expiredAt,
			Valid: true,
		},
	}

	err := s.sessionStorage.SaveSession(ctx, &session)

	if err != nil {
		return models.Session{}, err
	}

	return session, nil
}

func (s *SessionService) CheckSession(ctx *gin.Context, sessionId string) error {
	_, err := s.sessionStorage.GetSessionById(ctx, sessionId)

	if err != nil {
		return errors.New("Session with SessionId " + sessionId + " is not found")
	}

	return nil
}

func (s *SessionService) RemoveSession(ctx *gin.Context, sessionId string) error {
	err := s.sessionStorage.DeleteSession(ctx, sessionId)

	if err != nil {
		return err
	}

	return nil
}

func (s *SessionService) UpdateSessionExpiration(ctx *gin.Context, sessionId string) (models.Session, error) {
	session, err := s.sessionStorage.GetSessionById(ctx, sessionId)

	if err != nil {
		return models.Session{}, errors.New("Session with SessionId " + sessionId + " is not found")
	}

	/*if time.Now().Sub(session.ExpiredAt.Time).Minutes() > 3 {
		return models.Session{}, errors.New("token still has long lifetime")
	}*/

	session.ExpiredAt = sql.NullTime{
		Time:  time.Now().Add(time.Minute * SESSION_TTL),
		Valid: true,
	}

	err = s.sessionStorage.UpdateSessionExpiration(ctx, &session)

	if err != nil {
		return models.Session{}, err
	}

	return session, nil
}
