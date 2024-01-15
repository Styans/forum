package comment

import "forum/internal/models"

type CommentService struct {
	repo models.CommentRepo
}

func NewCommentService(repo models.CommentRepo) *CommentService {
	return &CommentService{repo}
}

func (s *CommentService) GetAllComments() ([]*models.CommentDTO, error)

func (s *CommentService) CreateComment(comment *models.CommentDTO) error

func (s *CommentService) UpdateComment(comment *models.CommentDTO) error

func (s *CommentService) DeleteComment(comment *models.CommentDTO) error
