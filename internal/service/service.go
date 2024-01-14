package service

import (
	"forum/internal/models"
	"forum/internal/repository"
	"forum/internal/service/user"
)

type Service struct {
	UserService models.UserService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService: user.NewUserService(repo.UserRepo),
	}
}
