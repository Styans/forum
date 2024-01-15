package post

import (
	"forum/internal/models"
)

type PostService struct {
	repo models.PostRepo
}

func NewPostService(repo models.PostRepo) *PostService {
	return &PostService{repo}
}

func (s *PostService) DeletePost(id int) error {
	return nil
}
func (s *PostService) CreatePost(post *models.CreatePostDTO) error {
	return nil
}

func (s *PostService) CreatePostWithImage(post *models.CreatePostDTO) error {
	return nil

}

func (s *PostService) UpdatePost(post *models.Post) error {
	return nil

}

func (s *PostService) GetPostsByAuthorID(author int) ([]*models.Post, error) {
	return nil, nil
}

func (s *PostService) GetAllPosts(offset, limit int) ([]*models.Post, error) {
	return nil, nil
}
