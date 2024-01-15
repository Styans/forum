package models

import "time"

type Session struct {
	UUID     string    `json:"uuid"`
	User_id  int       `json:"user_id"`
	ExpireAt time.Time `json:"expire_at"`
}

type SessionRepo interface {
	CreateSession(session *Session) error
	GetSessionByUserID(userUD int) (*Session, error)
	GetSessionByUUID(sessionID *Session) (*Session, error)
}

type SessionServise interface {
	CreateSession(userId int) (*Session, error)
}
