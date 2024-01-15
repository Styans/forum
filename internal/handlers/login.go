package handlers

import "net/http"

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user/login" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		h.templates.Render(w, r, "log.page.html", nil)
		return
	case http.MethodPost:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
