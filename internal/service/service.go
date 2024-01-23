package service

import (
	"forum/internal/models"
	"forum/internal/repository"
	"forum/internal/service/category"
	"forum/internal/service/comment"
	"forum/internal/service/post"
	"forum/internal/service/session"
	"forum/internal/service/user"
)

type Service struct {
	UserService     models.UserService
	PostService     models.PostService
	CommentService  models.CommentService
	SessionService  models.SessionServise
	CategoryService models.CategoryService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService:     user.NewUserService(repo.UserRepo),
		PostService:     post.NewPostService(repo.PostRepo),
		CommentService:  comment.NewCommentService(repo.CommentRepo),
		SessionService:  session.NewSessionService(repo.SessionRepo),
		CategoryService: category.NewCategoryService(repo.CategoryRepo),
	}
}
