package repository

import (
	"database/sql"
	"forum/internal/models"
	"forum/internal/repository/post"
	"forum/internal/repository/user"
)

type Repository struct {
	PostRepo models.PostRepo
	UserRepo models.UserRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		PostRepo: post.NewPostStorage(db),
		UserRepo: user.NewUserStorage(db),
	}
}
