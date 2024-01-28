package handlers

import (
	"net/http"

	"forum/internal/render"
)

func (h *Handler) showPostsByCategory(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/postscat" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	category := r.URL.Query().Get("category")

	posts, err := h.service.PostService.GetPostsByCategory(category)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.service.PostReactionService.GetAllPostReactionsByPostID(posts)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	categories, err := h.service.CategoryService.GetAllCategories()

	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	h.templates.Render(w, r, "home.page.html", &render.PageData{
		Topic:             category,
		Posts:             posts,
		Categories:        categories,
		AuthenticatedUser: h.getUserFromContext(r),
	})
}
