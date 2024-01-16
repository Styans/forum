package handlers

import (
	"context"
	"forum/internal/helpers/cookies"
	"forum/internal/models"
	"forum/internal/render"
	"forum/internal/service"
	"net/http"
	"time"
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
		// h.logger.PrintInfo("User is not authenticated")
		return nil
	}
	return user
}

func (h *Handler) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := cookies.GetCookie(r)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		session, err := h.service.SessionService.GetSessionByUUID(cookie.Value)
		if err != nil {
			cookies.DeleteCookie(w)
			next.ServeHTTP(w, r)
			return
		}

		if session.ExpireAt.Before(time.Now()) {
			cookies.DeleteCookie(w)
			next.ServeHTTP(w, r)
			return
		}

		user, err := h.service.UserService.GetUserByID(session.User_id)
		if err != nil {
			cookies.DeleteCookie(w)
			h.service.SessionService.DeleteSessionByUUID(cookie.Value)
			next.ServeHTTP(w, r)
		}

		ctx := context.WithValue(r.Context(), contextKeyUser, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
