package handlers

import (
	"html"
	"log"
	"net/http"
	"path"

	"github.com/adrg/frontmatter"
	"github.com/elkcityhazard/pc-building-company/content"
	"github.com/elkcityhazard/pc-building-company/internal/models"
)

func (hr *HandlerRepo) HandleGetServices(w http.ResponseWriter, r *http.Request) {
	contentFile, err := content.GetContentFS().Open(path.Join("content", "services.md"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer contentFile.Close()

	type Content struct {
		ID    int    `yaml:"id"`
		Name  string `yaml:"name"`
		URL   string `yaml:"url"`
		Text  string `yaml:"text"`
		Image string `yaml:"image"`
	}

	type Front struct {
		models.DefaultFrontMatter
		Content []Content `yaml:"content"`
	}

	var fm Front

	_, err = frontmatter.Parse(contentFile, &fm)

	hr.Cfg.Renderer.SetStringMapEntry("PageSubtitle", html.EscapeString("Home Builder In Traverse City, Leelanau County, and Grand Traverse County"))
	hr.Cfg.Renderer.SetDataMapEntry("FrontMatter", fm)
	hr.Cfg.Renderer.SetDataMapEntry("Content", hr.Cfg.MarkdownData["services.md"])
	err = hr.Cfg.Renderer.RenderTemplate(w, r, "services.gohtml")

	if err != nil {
		log.Fatalln(err)
	}
}
