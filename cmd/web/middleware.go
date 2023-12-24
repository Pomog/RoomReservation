package main

import (
	"net/http"
	helpers "udemyCourse1/internal/helper"

	"github.com/justinas/nosurf"
)

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// Auth 
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !helpers.IsAutenticated(r) {
			session.Put(r.Context(), "errpr", "Log in first")
			http.Redirect(w, r, "/user-login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})

}
