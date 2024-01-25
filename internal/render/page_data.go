package render

import (
	"forum/internal/models"
	"forum/pkg/forms"
)

type PageData struct {
	Topic             string
	Form              *forms.Form
	AuthenticatedUser *models.User
	Posts             []*models.Post
	Categories        []*models.Category
}
