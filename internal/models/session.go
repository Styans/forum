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
	GetSessionByUUID(sessionID string) (*Session, error)
	DeleteSessionByUUID(sessionID string) error
	// GetUserIdBySession(session *Session) (userid int, err error)
	GetUserIdBySession(session *Session) (int, error)
	// GetUserByID(id int) (user *User, err error)
}

type SessionServise interface {
	CreateSession(userId int) (*Session, error)
	DeleteSessionByUUID(uuid string) error
	GetUserIdBySession(session *Session) (int, error)
	GetSessionByUUID(uuid string) (*Session, error)
}
