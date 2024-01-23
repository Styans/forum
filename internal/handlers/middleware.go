package handlers

import (
	"context"
	"fmt"
	"forum/internal/helpers/cookies"
	"net/http"
	"time"
)


func (h *Handler) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := cookies.GetCookie(r)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		session, err := h.service.SessionService.GetSessionByUUID(cookie.Value)
		if err != nil {
			fmt.Println(err)
			cookies.DeleteCookie(w)
			next.ServeHTTP(w, r)
			return
		}

		if session.ExpireAt.Before(time.Now()) {
			cookies.DeleteCookie(w)
			next.ServeHTTP(w, r)
			return
		}

		user, err := h.service.UserService.GetUserByID(session.User_id)
		if err != nil {
			cookies.DeleteCookie(w)
			h.service.SessionService.DeleteSessionByUUID(cookie.Value)
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyUser, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
