package service

import (
	"forum/internal/models"
	"forum/internal/repository"

	"forum/internal/service/comment"
	"forum/internal/service/post"
	"forum/internal/service/session"
	"forum/internal/service/user"
)

type Service struct {
	UserService    models.UserService
	PostService    models.PostService
	CommentService models.CommentService
	SessionService models.SessionServise
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService:    user.NewUserService(repo.UserRepo),
		PostService:    post.NewPostService(repo.PostRepo),
		CommentService: comment.NewCommentService(repo.CommentRepo),
		SessionService: session.NewSessionService(repo.SessionRepo),
	}
}
