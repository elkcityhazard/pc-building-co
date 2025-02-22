package handlers

import (
	"bufio"
	"html"
	"html/template"
	"log"
	"net/http"

	"github.com/adrg/frontmatter"
	"github.com/elkcityhazard/pc-building-company/content"
	"github.com/elkcityhazard/pc-building-company/internal/models"
)

type HomeServices struct {
	Services []Service `json:"services"`
}

type Service struct {
	Name        string        `json:"name"`
	Url         string        `json:"url"`
	Description string        `json:"description"`
	Icon        template.HTML `json:"icon"`
}

func (hr *HandlerRepo) GetHomeHandler(w http.ResponseWriter, r *http.Request) {

	var fm models.DefaultFrontMatter

	contentFiles := content.GetContentFS()

	f, err := contentFiles.Open("content/home.md")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = frontmatter.Parse(bufio.NewReader(f), &fm)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hr.Cfg.Renderer.SetStringMapEntry("PageTitle", hr.Cfg.WebsiteName)
	hr.Cfg.Renderer.SetStringMapEntry("PageSubtitle", html.EscapeString("Home Builder In Traverse City, Leelanau County, and Grand Traverse County"))
	hr.Cfg.Renderer.SetDataMapEntry("FrontMatter", fm)
	err = hr.Cfg.Renderer.RenderTemplate(w, r, "home.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

}
