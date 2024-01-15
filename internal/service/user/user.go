package user

import (
	"forum/internal/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo models.UserRepo
}

func NewUserService(repo models.UserRepo) *UserService {
	return &UserService{repo}
}

func (u *UserService) CreateUser(userDTO *models.CreateUserDTO) error {
	hashedPW, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username:  userDTO.Username,
		Email:     userDTO.Email,
		HashedPW:  string(hashedPW),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = u.repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
