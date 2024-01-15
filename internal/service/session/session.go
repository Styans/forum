package session

import (
	"forum/internal/models"
	"time"

	"github.com/gofrs/uuid"
)

type SessionService struct {
	repo models.SessionRepo
}

func NewSessionService(repo models.SessionRepo) *SessionService {
	return &SessionService{repo}
}

func (s *SessionService) CreateSession(userId int) (*models.Session, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	session := &models.Session{
		UUID:     uuid.String(),
		User_id:  userId,
		ExpireAt: time.Now().Add(time.Hour),
	}
	err = s.repo.CreateSession(session)
	if err != nil {
		return nil, err
	}
	return session, nil
}
