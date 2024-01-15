package post

import (
	"database/sql"
	"forum/internal/models"
)

type PostStorage struct {
	db *sql.DB
}

func NewPostStorage(db *sql.DB) *PostStorage {
	return &PostStorage{db: db}
}

func (s *PostStorage) CreatePost(post *models.Post) error {
	return nil
}

func (s *PostStorage) DeletePost(post *models.Post) error {
	return nil
}

func (s *PostStorage) UpdatePost(post *models.Post) error {
	return nil
}

func (s *PostStorage) GetAllPosts(offset, limit int) ([]*models.Post, error) {
	return nil, nil
}

