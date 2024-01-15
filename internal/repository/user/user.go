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

//registration user ===================================================
func (s *UserStorage) CreateUser(user *models.User) error {
	_, err := s.db.Exec("INSERT INTO user (username, hashed_pw, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		user.Username,
		user.HashedPW,
		user.Email,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		switch err.Error() {
		case "UNIQUE constraint failed: user.email":
			return models.ErrDuplicateEmail
		case "UNIQUE constraint failed: user.username":
			return models.ErrDuplicateUsername
		default:
			return err
		}
	}

	return nil
}

// for authentification user=============================================
func (s *UserStorage) GetUserByUsername(email string) (*models.User, error) {
	return nil, nil
}
func (s *UserStorage) GetUserByEmail(email string) (*models.User, error) {
	return nil, nil
}

// for update user datas=========================================
func (s *UserStorage) UpdateUser(user *models.User) error {
	return nil
}

func (s *UserStorage) GetAllUsers() ([]*models.User, error) {
	return nil, nil
}

// administration.UsersFuncs =====================================
func (s *UserStorage) DeleteUser(user *models.User) error {
	return nil
}
