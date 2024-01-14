package user

import (
	"database/sql"
	"forum/internal/models"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (s *UserStorage) CreateUser(user *models.User) error {

	_, err := s.db.Exec("INSERT INTO user (username, hashed_pw, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		user.Username,
		user.HashedPW,
		user.Email,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
