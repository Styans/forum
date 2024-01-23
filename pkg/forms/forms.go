package forms

import (
	"forum/internal/models"
	"mime/multipart"
	"net/url"
	"strings"
)

const AllowedTypes = "image/jpeg,image/png,image/gif,image/jpg,image/webp"

func IsImg(fileType string) bool {
	return strings.Contains(AllowedTypes, fileType)
}

type errors map[string][]string

func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}

type Form struct {
	url.Values
	Errors     errors
	Categories []*models.Category
	FileHeader *multipart.FileHeader
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
		[]*models.Category{},
		nil,
	}
}
