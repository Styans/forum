package handlers

import (
	"forum/internal/models"
	"forum/pkg/forms"
	"net/http"
	"strings"
)

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post/create" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Incorrect Method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10); err != nil {
		http.Error(w, "Invalid POST request", http.StatusInternalServerError)
		return
	}
	autor := h.getUserFromContext(r)
	post := &models.CreatePostDTO{
		Title:      r.PostFormValue("title"),
		Content:    r.PostFormValue("content"),
		Author:     autor.ID,
		AuthorName: autor.Username,
		// Categories: categories,
	}

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileType := fileHeader.Header.Get("Content-Type")

	if !forms.IsImg(fileType) {
		http.Error(w, "nil", http.StatusInternalServerError)
		return
	}

	post.ImageFile = file
	defer file.Close()
	datas := r.PostFormValue("category")
	tempD := strings.Split(datas, ",")
	for i, v := range tempD {
		tempD[i] = strings.TrimSpace(v)
	}
	categories := make([]*models.Category, 0, len(tempD))
	for _, name := range tempD {
		c, err := h.service.CategoryService.GetCategoryByName(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			categories = append(categories, c)
		}
	}

	post.Categories = append(post.Categories, categories...)
	_, err = h.service.PostService.CreatePostWithImage(post)

	if err != nil {

		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
