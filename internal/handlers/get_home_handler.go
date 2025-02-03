package handlers

import (
	"html"
	"html/template"
	"log"
	"net/http"
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
	hr.Cfg.Renderer.SetStringMapEntry("PageSubtitle", html.EscapeString("Home Builder In Traverse City, Leelanau County, and Grand Traverse County"))
	err := hr.Cfg.Renderer.RenderTemplate(w, r, "home.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

}
