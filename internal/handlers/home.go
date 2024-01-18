package handlers

import (
	"fmt"
	"forum/internal/render"
	"net/http"
	"strconv"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}
	posts, err := h.service.PostService.GetAllPosts(offset, limit)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	h.templates.Render(w, r, "home.page.html", &render.PageData{
		Posts:             posts,
		AuthenticatedUser: h.getUserFromContext(r),
	})
}
