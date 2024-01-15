package handlers

import (
	"context"
	"net/http"
)

func (h *Handler) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	// add a css file to route
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/user/register", h.register)
	mux.HandleFunc("/user/login", h.login)

	mux.HandleFunc("/", isAuth(rateLimit(h.home)))

	return mux
}

func rateLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// here
		next.ServeHTTP(w, r)
	}
}

func isAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token from cookie

		token := ""
		// if token not exist
		// w.Write([]byte)
		// return

		_ = token
		// get user from token
		user := "user"
		_ = user

		ctx := r.Context()
		ctx = context.WithValue(ctx, "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
