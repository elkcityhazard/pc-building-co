package main

import (
	"fmt"
	"html"
	"net/http"

	"github.com/elkcityhazard/pc-building-company/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.Recoverer)
	mux.Use(preventStaticDirBrowsing)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", handlers.Repo.GetHomeHandler)
	mux.HandleFunc("/about", handlers.Repo.GetAboutHandler)
	mux.HandleFunc("/contact", handlers.Repo.GetContactHandler)
	mux.HandleFunc("/services", handlers.Repo.HandleGetServices)
	mux.HandleFunc("/services/{service}", handlers.Repo.HandleGetService)

	mux.HandleFunc("/robots.txt", handlers.Repo.HandleGetRobots)
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
