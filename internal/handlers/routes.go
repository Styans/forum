package handlers

import (
	"net/http"
)

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	// add a css file to route
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/user/register", h.register)
	mux.HandleFunc("/user/login", h.login)
	mux.HandleFunc("/user/logout", h.logout)
	mux.HandleFunc("/post/create", h.createPost)
	mux.HandleFunc("/", h.home)
	mux.HandleFunc("/posts", h.GetPosts)

	return h.authenticate(mux)
}

func rateLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// here
		next.ServeHTTP(w, r)
	}
}
