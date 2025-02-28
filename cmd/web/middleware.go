package main

import (
	"net/http"
	"strings"

	"github.com/justinas/nosurf"
)

var (
	pathToServices   string = "./static/data/services.json"
	markdownBasePath string = "./content"
)

func preventStaticDirBrowsing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/static/") {
			if strings.HasSuffix(r.URL.Path, "/") {
				http.NotFound(w, r)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func staticCacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request URL path has the "/static/" prefix
		if strings.HasPrefix(r.URL.Path, "/static/") {
			// Set Cache-Control header for 1 year
			w.Header().Set("Cache-Control", "public, max-age=31536000")
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	// new no surf
	csrfHandler := nosurf.New(next)

	// cookie config
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.IsProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func AddCSRFToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.Renderer.AddCSRFToken(nosurf.Token(r))
		next.ServeHTTP(w, r)
	})
}
