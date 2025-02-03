package handlers

import (
	"bufio"
	"html"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/adrg/frontmatter"
	"github.com/elkcityhazard/pc-building-company/content"
	"github.com/elkcityhazard/pc-building-company/internal/models"
)

func (hr *HandlerRepo) GetAboutHandler(w http.ResponseWriter, r *http.Request) {

	aboutMD, err := content.GetContentFS().Open(filepath.Join("content", "about.md"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer aboutMD.Close()

	type FrontMatter struct {
		models.DefaultFrontMatter
	}

	var fm FrontMatter

	_, err = frontmatter.Parse(bufio.NewReader(aboutMD), &fm)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hr.Cfg.Renderer.SetStringMapEntry("PageSubtitle", html.EscapeString("Home Builder In Traverse City, Leelanau County, and Grand Traverse County"))
	hr.Cfg.Renderer.SetDataMapEntry("FrontMatter", fm)
	hr.Cfg.Renderer.SetDataMapEntry("Content", template.HTML(hr.Cfg.MarkdownData["about.md"]))
	err = hr.Cfg.Renderer.RenderTemplate(w, r, "about.gohtml")

	if err != nil {
		log.Fatalln(err)
	}
}
