package render

import "forum/internal/models"

type PageData struct {
	AuthenticatedUser *models.User
	Posts             []*models.Post
}
