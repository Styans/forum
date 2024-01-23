package post

import (
	"forum/internal/models"
	"io/ioutil"
	"time"

	"github.com/gofrs/uuid"
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

func (p *PostService) CreatePost(postDTO *models.CreatePostDTO) (int, error) {
	post := &models.Post{
		Title:      postDTO.Title,
		Content:    postDTO.Content,
		AuthorID:   postDTO.Author,
		AuthorName: postDTO.AuthorName,
		Categories: postDTO.Categories,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	return p.repo.CreatePost(post)
}

func (p *PostService) CreatePostWithImage(postDTO *models.CreatePostDTO) (int, error) {
	if postDTO.ImageFile == nil {
		return p.CreatePost(postDTO)
	}

	data, err := ioutil.ReadAll(postDTO.ImageFile)
	if err != nil {
		return 0, err
	}

	fileName, err := uuid.NewV4()
	if err != nil {
		return 0, err
	}
	filePath := "ui/static/img/" + fileName.String()

	post := &models.Post{
		Title:      postDTO.Title,
		Content:    postDTO.Content,
		AuthorID:   postDTO.Author,
		AuthorName: postDTO.AuthorName,
		Categories: postDTO.Categories,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		ImagePath:  filePath,
	}

	id, err := p.repo.CreatePostWithImage(post)
	if err != nil {
		return 0, err
	}

	err = ioutil.WriteFile(filePath, data, 0o666)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *PostService) UpdatePost(post *models.Post) error {
	return nil
}

func (s *PostService) GetPostsByAuthorID(author int) ([]*models.Post, error) {
	return nil, nil
}

func (s *PostService) GetAllPosts(offset, limit int) ([]*models.Post, error) {
	return s.repo.GetAllPosts(offset, limit)
}
