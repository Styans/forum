package handlers

import (
	"fmt"
	"forum/internal/render"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) showPost(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/post/") {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	pathID := r.URL.Path[len("/post/"):]
	id, err := strconv.Atoi(pathID)
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	post, err := h.service.PostService.GetPostByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	comments, err := h.service.CommentService.GetAllByPostID(post.ID)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, comment := range comments {
		comment.Likes, comment.Dislikes, err = h.service.CommentReactionService.GetLikesAndDislikes(comment.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	err = h.service.PostReactionService.PutReactionsToPost(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	categories, err := h.service.CategoryService.GetAllCategories()

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	h.templates.Render(w, r, "post.page.html", &render.PageData{
		Topic:             post.AuthorName,
		Post:              post,
		Comments:          comments,
		Categories:        categories,
		AuthenticatedUser: h.getUserFromContext(r),
	})
}
