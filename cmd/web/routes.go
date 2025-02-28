package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"path/filepath"

	"github.com/elkcityhazard/pc-building-company/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.Recoverer)
	mux.Use(preventStaticDirBrowsing)
	mux.Use(staticCacheMiddleware)
	mux.Use(NoSurf)
	mux.Use(AddCSRFToken)
	mux.Use(middleware.Compress(5))

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", handlers.Repo.GetHomeHandler)
	mux.HandleFunc("/about", handlers.Repo.GetAboutHandler)
	mux.HandleFunc("/contact", handlers.Repo.GetContactHandler)
	mux.HandleFunc("/services", handlers.Repo.HandleGetServices)
	// mux.HandleFunc("/services/{service}", handlers.Repo.HandleGetService)
	mux.Post("/contact", handlers.Repo.PostContactHandler)

	mux.Handle("/robots.txt", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cwd, err := os.Getwd()
		if err != nil {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, filepath.Join(cwd, "robots.txt"))

	}))
	mux.Handle("/ads.txt", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cwd, err := os.Getwd()

		if err != nil {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, filepath.Join(cwd, "ads.txt"))
	}))
	mux.NotFound(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.Renderer.SetStringMapEntry("PageSubtitle", html.EscapeString("Home Builder In Traverse City, Leelanau County, and Grand Traverse County"))
		app.Renderer.SetStringMapEntry("ContactLink", "/contact")
		app.Renderer.SetStringMapEntry("PhoneNumber", "+12313576340")

		err := app.Renderer.RenderTemplate(w, r, "404.gohtml")

		if err != nil {
			fmt.Fprint(w, "not found")
		}
	}))

	return mux
}
