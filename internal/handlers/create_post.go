package handlers

import "net/http"

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post/create" {
		
		return 
	}
}
