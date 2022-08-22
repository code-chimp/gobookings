package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// NoSurf add XSRF token cookie
func NoSurf(next http.Handler) http.Handler {
	xsrfHandler := nosurf.New(next)

	xsrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return xsrfHandler
}

// SessionLoad load and save session on each request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
