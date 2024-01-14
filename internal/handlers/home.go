package handlers

import (
	"forum/internal/render"
	"net/http"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	h.templates.Render(w, r, "home.page.html", &render.PageData{
		AuthenticatedUser: h.getUserFromContext(r),
	})
}
