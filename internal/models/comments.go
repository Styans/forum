package models

type Comment struct {
}

type CommentDTO struct {
}
type CommentRepo interface {
	GetAllComments() ([]*Comment, error)
	CreateComment(comment *Comment) error
	UpdateComment(comment *Comment) error
	DeleteComment(id int) error
}

type CommentService interface {
	GetAllComments() ([]*CommentDTO, error)
	CreateComment(comment *CommentDTO) error
	UpdateComment(comment *CommentDTO) error
	DeleteComment(comment *CommentDTO) error
}
