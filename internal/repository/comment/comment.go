package comment

import (
	"database/sql"
	"forum/internal/models"
)

type CommentStorage struct {
	db *sql.DB
}

func NewCommentStorage(db *sql.DB) *CommentStorage {
	return &CommentStorage{db: db}
}

func (s *CommentStorage) GetAllComments() ([]*models.Comment, error) {
	return nil, nil
}

// administration.UsersFuncs =====================================

func (s *CommentStorage) CreateComment(comment *models.Comment) error {
	return nil
}

func (s *CommentStorage) UpdateComment(comment *models.Comment) error {
	return nil
}
func (s *CommentStorage) DeleteComment(id int) error {
	return nil
}
