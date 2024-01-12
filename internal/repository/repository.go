package repository

import (
	"database/sql"
	"forum/internal/models"
	"forum/internal/repository/post"
)

type Repository struct {
	PostRepo models.PostRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		PostRepo: post.NewPostStorage(db),
	}
}
