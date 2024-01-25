package handlers

import (
	"forum/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) reactionPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post/reaction" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Incorrect Method", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Incorrect Method", http.StatusMethodNotAllowed)
		return
	}

	userID := h.getUserFromContext(r)
	postID, err := strconv.Atoi(r.FormValue("post_id"))
	if err != nil {
		return
	}
	data, err := strconv.Atoi(r.FormValue("status"))
	if err != nil {
		return
	}
	var status bool
	switch data {
	case 1:
		status = true
	case 0:
		status = false
	default:

		// fmt.Println(data)
		http.Error(w, "Incorrect Method", http.StatusMethodNotAllowed)
		return
	}

	reaction := &models.PostReactionDTO{
		UserID: userID.ID,
		PostID: postID,
		Status: status,
	}
	err = h.service.PostReactionService.CreatePostReaction(reaction)
	if err != nil {
		http.Error(w, "Incorrect Method", http.StatusMethodNotAllowed)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
