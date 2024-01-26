package handlers

import (
	"fmt"
	"forum/internal/models"
	"log"
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
		log.Printf("Error parsing form: %v", err)
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	userID := h.getUserFromContext(r)

	postID, err := strconv.Atoi(r.FormValue("post_id"))
	if err != nil {
		log.Printf("Error converting post_id: %v", err)
		http.Error(w, "Invalid post_id", http.StatusBadRequest)
		return
	}

	status, err := strconv.Atoi(r.FormValue("status"))
	if err != nil {
		log.Printf("Error converting status: %v", err)
		http.Error(w, "Invalid status", http.StatusBadRequest)
		return
	}

	switch status {
	case 1:
		// Status is true
	case 0:
		// Status is false
	default:
		log.Println("Invalid status value")
		http.Error(w, "Invalid status value", http.StatusBadRequest)
		return
	}

	reaction := &models.PostReactionDTO{
		UserID: userID.ID,
		PostID: postID,
		Status: status == 1,
	}

	err = h.service.PostReactionService.CreatePostReaction(reaction)
	if err != nil {
		log.Printf("Error creating post reaction: %v", err)
		http.Error(w, "Error creating post reaction", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/post/%d", postID), http.StatusFound)
}
