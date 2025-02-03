package handlers

import (
	"bufio"
	"html"
	"log"
	"net/http"
	"path/filepath"

	"github.com/adrg/frontmatter"
	"github.com/elkcityhazard/pc-building-company/content"
	"github.com/elkcityhazard/pc-building-company/internal/models"
)

func (hr *HandlerRepo) GetContactHandler(w http.ResponseWriter, r *http.Request) {
	type Matter struct {
		models.DefaultFrontMatter
	}

	var matter Matter

	contentFS := content.GetContentFS()

	toRead, err := contentFS.Open(filepath.Join("content", "contact.md"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = frontmatter.Parse(bufio.NewReader(toRead), &matter)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hr.Cfg.Renderer.SetStringMapEntry("PageSubtitle", html.EscapeString("Home Builder In Traverse City, Leelanau County, and Grand Traverse County"))
	hr.Cfg.Renderer.SetDataMapEntry("FrontMatter", matter)
	hr.Cfg.Renderer.SetDataMapEntry("Content", hr.Cfg.MarkdownData["contact.md"])
	err = hr.Cfg.Renderer.RenderTemplate(w, r, "contact.gohtml")

	if err != nil {
		log.Fatalln(err)
	}
}
