package handlers

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"path/filepath"

	"github.com/adrg/frontmatter"
	"github.com/elkcityhazard/pc-building-company/content"
	"github.com/elkcityhazard/pc-building-company/internal/models"
	"github.com/go-chi/chi/v5"
)

func (hr *HandlerRepo) HandleGetService(w http.ResponseWriter, r *http.Request) {
	serviceName := chi.URLParam(r, "service")

	fname := fmt.Sprintf("%s.md", serviceName)

	contentFile, err := content.GetContentFS().Open(filepath.Join("content", fname))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer contentFile.Close()

	var fm models.DefaultFrontMatter
	_, err = frontmatter.Parse(contentFile, &fm)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hr.Cfg.Renderer.SetStringMapEntry("PageSubtitle", html.EscapeString("Home Builder In Traverse City, Leelanau County, and Grand Traverse County"))
	hr.Cfg.Renderer.SetDataMapEntry("FrontMatter", fm)
	hr.Cfg.Renderer.SetDataMapEntry("Content", hr.Cfg.MarkdownData[fmt.Sprintf("%s", fname)])
	err = hr.Cfg.Renderer.RenderTemplate(w, r, "service.gohtml")

	if err != nil {
		log.Fatalln(err)
	}
}
