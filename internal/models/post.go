package models

import (
	"mime/multipart"
	"time"
)

type Post struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	AuthorID   int       `json:"author_id"`
	AuthorName string    `json:"authorname"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Image      []byte    `json:"image"`
}

type CreatePostDTO struct {
	Title      string         `json:"title"`
	Content    string         `json:"content"`
	Author     int            `json:"author"`
	AuthorName string         `json:"authorname"`
	ImageFile  multipart.File `json:"imagefile"`
}

type UpdatePostDTO struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Category_id int    `json:"category_id"`
}

type DeletePostDTO struct {
	ID int `json:"id"`
}

type PostService interface {
	CreatePost(p *CreatePostDTO) error
	CreatePostWithImage(post *CreatePostDTO) error
	GetAllPosts(offset, limit int) ([]*Post, error)
	GetPostsByAuthorID(author int) ([]*Post, error)
	UpdatePost(post *Post) error
	DeletePost(id int) error
}

type PostRepo interface {
	GetAllPosts(offset, limit int) ([]*Post, error)
	CreatePost(p *Post) error
}
