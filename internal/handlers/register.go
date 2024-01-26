package handlers

import (
	"forum/internal/models"
	"forum/internal/render"
	"forum/pkg/forms"
	"net/http"
)

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
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

		form := forms.New(r.PostForm)
		form.Required("username", "email", "password")
		form.MaxLength("username", 50)
		form.MaxLength("email", 50)
		form.MatchesPattern("email", forms.EmailRX)
		form.MaxLength("password", 50)
		form.MinLength("password", 8)

		if !form.Valid() {
			form.Errors.Add("generic", "Invalid credentials")
			w.WriteHeader(http.StatusBadRequest)
			h.templates.Render(w, r, "reg.page.html", &render.PageData{
				Form: form,
			})
			return
		}

		req := &models.CreateUserDTO{
			Email:    r.FormValue("email"),
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}

		err = h.service.UserService.CreateUser(req)

		if err != nil {
			switch err {
			case models.ErrDuplicateEmail:
				form.Errors.Add("email", "Email already in use")
				w.WriteHeader(http.StatusBadRequest)
				h.templates.Render(w, r, "reg.page.html", &render.PageData{
					Form: form,
				})
				return
			case models.ErrDuplicateUsername:
				form.Errors.Add("username", "Username already in use")
				w.WriteHeader(http.StatusBadRequest)
				h.templates.Render(w, r, "reg.page.html", &render.PageData{
					Form: form,
				})
				return
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	case http.MethodGet:
		h.templates.Render(w, r, "reg.page.html", &render.PageData{
			Form:              forms.New(nil),
			AuthenticatedUser: h.getUserFromContext(r),
		})
		return
	default:
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	return
}
