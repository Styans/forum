package handlers

import (
	"forum/internal/models"
	"net/http"
)

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post/create" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "incorrect Method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid POST request", http.StatusInternalServerError)
		return
	}

	autor := h.getUserFromContext(r)
	post := &models.CreatePostDTO{
		Author:     autor.ID,
		AuthorName: autor.Username,
		Title:      r.FormValue("title"),
		Content:    r.FormValue("content"),
	}
	// fmt.Println(r.FormValue("title"))
	// fmt.Println(r.FormValue("content"))
	// fmt.Println(post)
	// file, file_header, err := r.FormFile("imagefile")
	// filetype := file_header.Header.Get("Content-type")

	// switch {
	// case err != nil:
	// 	http.Error(w, err.Error(), http.StatusMethodNotAllowed)
	// 	return
	// case !forms.IsImg(filetype):
	// 	http.Error(w, err.Error(), http.StatusMethodNotAllowed)
	// 	return
	// }

	// post.ImageFile = file
	err = h.service.PostService.CreatePost(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
