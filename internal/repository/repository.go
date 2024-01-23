package repository

import (
	"database/sql"
	"forum/internal/models"

	// "forum/internal/repository/category"
	"forum/internal/repository/category"
	"forum/internal/repository/comment"
	"forum/internal/repository/post"
	"forum/internal/repository/session"
	"forum/internal/repository/user"
)

type Repository struct {
	CommentRepo  models.CommentRepo
	PostRepo     models.PostRepo
	UserRepo     models.UserRepo
	SessionRepo  models.SessionRepo
	CategoryRepo models.CategoryRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		PostRepo:     post.NewPostStorage(db),
		UserRepo:     user.NewUserStorage(db),
		CommentRepo:  comment.NewCommentStorage(db),
		SessionRepo:  session.NewSessionStorage(db),
		CategoryRepo: category.NewCategoryStorage(db),
	}
}
