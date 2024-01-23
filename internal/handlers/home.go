package handlers

import (
	"fmt"
	"forum/internal/render"
	"forum/pkg/forms"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	form := forms.New(r.PostForm)
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}
	// limit := 10
	// offset := 0
	posts, err := h.service.PostService.GetAllPosts(offset, limit)
	if err != nil {
		log.Println(err)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	categories, err := h.service.CategoryService.GetAllCategories()

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	h.templates.Render(w, r, "home.page.html", &render.PageData{
		Categories:        categories,
		Form:              form,
		Posts:             posts,
		AuthenticatedUser: h.getUserFromContext(r),
	})
}

func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}
	// limit := 10
	// offset := 0
	posts, err := h.service.PostService.GetAllPosts(offset, limit)
	if err != nil {
		log.Println(err)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	h.templates.Render(w, r, "posts.page.html", &render.PageData{Posts: posts})
}
