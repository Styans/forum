package handlers

import (
	"forum/internal/render"
	"net/http"
)

func (h *Handler) myposts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/myposts" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user := h.getUserFromContext(r)
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	posts, err := h.service.PostService.GetPostsByAuthorID(user.ID)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
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
		Topic:             "LikedPosts",
		Categories:        categories,
		Posts:             posts,
		AuthenticatedUser: user,
	})

}
