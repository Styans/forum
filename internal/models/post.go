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

	Categories []*Category `json:"category_id"`
	Comments   []*Comment  `json:"comments"`
	Likes      int         `json:"likes"`
	Dislikes   int         `json:"dislikes"`
	ImagePath  string      `json:"image_path"`
	Image      []byte      `json:"image"`
}

type CreatePostDTO struct {
	Title      string         `json:"title"`
	Content    string         `json:"content"`
	Categories []*Category    `json:"category_id"`
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
	CreatePost(p *CreatePostDTO) (int, error)
	CreatePostWithImage(post *CreatePostDTO) (int, error)

	GetAllPosts(offset, limit int) ([]*Post, error)
	GetPostsByAuthorID(author int) ([]*Post, error)
	UpdatePost(post *Post) error
	DeletePost(id int) error

}

type PostRepo interface {
	CreatePost(p *Post) (int, error)
	CreatePostWithImage(post *Post) (int, error)
	GetAllPosts(offset, limit int) ([]*Post, error)

}
