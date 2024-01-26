package handlers

import (
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

	// comments, err := h.service.CommentService.GetAllByPostID(post.ID)
	// if err != nil {

	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// for _, comment := range comments {
	// 	// comment.Likes, comment.Dislikes, err = h.service.CommentVoteService.GetLikesAndDislikes(comment.ID)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// }

	err = h.service.PostReactionService.PutReactionsToPost(post)
	if err != nil {
		// h.logger.PrintError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.templates.Render(w, r, "post.page.html", &render.PageData{
		Post: post,
		// Comments:          comments,
		AuthenticatedUser: h.getUserFromContext(r),
	})
}
