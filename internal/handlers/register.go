package handlers

import (
	"forum/internal/models"
	"net/http"
)

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user/register" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodPost:

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		req := &models.CreateUserDTO{
			Email:    r.FormValue("email"),
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}

		err = h.service.UserService.CreateUser(req)

		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}

		http.Redirect(w, r, "/user/login", http.StatusSeeOther)
	case http.MethodGet:
		h.templates.Render(w, r, "reg.page.html", nil)
		return
	default:
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	return
}
