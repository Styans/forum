package handlers

import "net/http"

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user/login" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
