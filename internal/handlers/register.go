package handlers

import (
	"fmt"
	"forum/internal/models"
	"forum/internal/render"
	"net/http"
)

func (h *Handler) registers(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user/signup" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	h.templates.Render(w, r, "reg.page.html", &render.PageData{

		AuthenticatedUser: h.getUserFromContext(r),
	})
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user/register" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := &models.CreateUserDTO{
		Username: r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.PostFormValue("password"),
	}

	fmt.Println(err, "ASDSD\n\n\n", user)
	err = h.service.UserService.CreateUser(user)

	if err != nil {
		fmt.Println(err, "ASDSD\n\n\n")
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)

	return
}
