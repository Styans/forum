package post

import (
	"context"
	"database/sql"
	"forum/internal/models"
	"time"
)

type PostStorage struct {
	db *sql.DB
}

func NewPostStorage(db *sql.DB) *PostStorage {
	return &PostStorage{db: db}
}

func (s *PostStorage) CreatePost(p *models.Post) (int, error) {
	query := `INSERT INTO posts (title, content, author_id, authorname, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, created_at, updated_at`

	args := []interface{}{p.Title, p.Content, p.AuthorID, p.AuthorName, p.CreatedAt, p.UpdatedAt}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, args...).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return 0, err
		// return 0, err
	}

	for _, category := range p.Categories {

		query = `INSERT INTO PostCategories (post_id, category_name) VALUES ($1, $2)`
		_, err = s.db.ExecContext(ctx, query, p.ID, category.Name)
		if err != nil {
			return 0, err
		}
	}

	return p.ID, nil
}

func (s *PostStorage) CreatePostWithImage(p *models.Post) (int, error) {
	query := `INSERT INTO posts (title, content, author_id, authorname, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING id, created_at, updated_at`

	args := []interface{}{p.Title, p.Content, p.AuthorID, p.AuthorName, p.CreatedAt, p.UpdatedAt, p.ImagePath}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, args...).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return 0, err
	}

	// Create post_categories entries
	// fmt.Println(p)

	for _, category := range p.Categories {
		// fmt.Println("================================================")

		// fmt.Println(category)
		// fmt.Println(p)
		// fmt.Println("================================================")
		query = `INSERT INTO PostCategories (post_id, category_name) VALUES ($1, $2)`
		_, err = s.db.ExecContext(ctx, query, p.ID, category.Name)
		if err != nil {
			return 0, err
		}
	}

	// Create image entry
	query = `INSERT INTO images (post_id, image_path) VALUES ($1, $2)`
	_, err = s.db.ExecContext(ctx, query, p.ID, p.ImagePath)
	if err != nil {
		return 0, err
	}

	return p.ID, nil
}

func (s *PostStorage) DeletePost(post *models.Post) error {
	return nil
}

func (s *PostStorage) UpdatePost(post *models.Post) error {
	return nil
}

func (s *PostStorage) GetAllPosts(offset, limit int) ([]*models.Post, error) {
	query := `SELECT * FROM posts ORDER BY id DESC LIMIT $1 OFFSET $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post

	for rows.Next() {
		post := models.Post{}

		err := rows.Scan(&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.AuthorName,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
