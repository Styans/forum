package post

import (
	"forum/internal/models"
	"time"
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

func (s *PostService) CreatePost(postDTO *models.CreatePostDTO) error {
	post := &models.Post{
		Title:      postDTO.Title,
		Content:    postDTO.Content,
		AuthorID:   postDTO.Author,
		AuthorName: postDTO.AuthorName,
		// Categories: postDTO.Categories,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		// Version:    1,
	}
	return s.repo.CreatePost(post)
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
	return s.repo.GetAllPosts(offset,limit)
}
