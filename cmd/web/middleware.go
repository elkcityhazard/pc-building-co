package main

import (
	"net/http"
	"strings"
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
