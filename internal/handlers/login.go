package handlers

import (
	"forum/internal/helpers/cookies"
	"forum/internal/models"
	"net/http"
	"time"
)

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
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid POST request", http.StatusInternalServerError)
			return
		}
		req := &models.LoginUserDTO{
			Email:    r.FormValue("email"),
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		user_id, err := h.service.UserService.LoginUser(req)
		if err != nil {
			http.Error(w, "User not found", http.StatusBadGateway)
			return
		}
		session, err := h.service.SessionService.CreateSession(user_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cookies.SetCookie(w, session.UUID, int(time.Until(session.ExpireAt).Seconds()))
	default:
		http.Error(w, "incorrect Method", http.StatusMethodNotAllowed)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
