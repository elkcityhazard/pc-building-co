package handlers

import (
	"bufio"
	"html"
	"log"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/adrg/frontmatter"
	"github.com/elkcityhazard/pc-building-company/content"
	"github.com/elkcityhazard/pc-building-company/internal/models"
)

func (hr *HandlerRepo) GetContactHandler(w http.ResponseWriter, r *http.Request) {
	type Matter struct {
		models.DefaultFrontMatter
	}

	var matter models.DefaultFrontMatter

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

	hr.Cfg.Renderer.SetStringMapEntry("PageTitle", html.EscapeString(matter.Title))
	hr.Cfg.Renderer.SetStringMapEntry("PageSubtitle", html.EscapeString(matter.Description))
	hr.Cfg.Renderer.SetDataMapEntry("FrontMatter", matter)
	hr.Cfg.Renderer.SetDataMapEntry("Content", hr.Cfg.MarkdownData["contact.md"])
	hr.Cfg.Renderer.SetDataMapEntry("FormErrors", make(map[string][]string)) // init empty
	hr.Cfg.Renderer.SetDataMapEntry("Form", &url.Values{})                   // init empty
	err = hr.Cfg.Renderer.RenderTemplate(w, r, "contact.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
}
