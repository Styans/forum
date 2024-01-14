package handlers

import (
	"forum/internal/models"
	"forum/internal/render"
	"forum/internal/service"
	"net/http"
)

type Handler struct {
	service   *service.Service
	templates render.TemplatesHTML
}

func NewHandler(service *service.Service, tmlp render.TemplatesHTML) *Handler {
	return &Handler{
		service:   service,
		templates: tmlp,
	}
}

type contextKey string

var contextKeyUser = contextKey("user")

func (h *Handler) getUserFromContext(r *http.Request) *models.User {
	user, ok := r.Context().Value(contextKeyUser).(*models.User)
	if !ok {
		return nil
	}
	return user
}
