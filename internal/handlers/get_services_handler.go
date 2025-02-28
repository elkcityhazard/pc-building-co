package handlers

import (
	"bufio"
	"html"
	"log"
	"net/http"
	"path"

	"github.com/adrg/frontmatter"
	"github.com/elkcityhazard/pc-building-company/content"
)

func (hr *HandlerRepo) HandleGetServices(w http.ResponseWriter, r *http.Request) {
	contentFile, err := content.GetContentFS().Open(path.Join("content", "services.md"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer contentFile.Close()

	type Content struct {
		ID       int      `yaml:"id"`
		Name     string   `yaml:"name"`
		URL      string   `yaml:"url"`
		Services []string `yaml:"services"`
		Text     string   `yaml:"text"`
		Image    string   `yaml:"image"`
	}

	type DefaultWithContent struct {
		Title           string    `yaml:"title"`
		Description     string    `yaml:"description"`
		URL             string    `yaml:"url"`
		MetaImage       string    `yaml:"metaImage"`
		MetaImageAlt    string    `yaml:"metaImageAlt"`
		Author          string    `yaml:"author"`
		PublishDate     string    `yaml:"publishDate"`
		UpdateDate      string    `yaml:"updateDate"`
		BgImage         string    `yaml:"bgImage"`
		TwitterUsername string    `yaml:"twitterUsername"`
		Content         []Content `yaml:"content"`
	}

	var fm DefaultWithContent

	_, err = frontmatter.Parse(bufio.NewReader(contentFile), &fm)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hr.Cfg.Renderer.SetStringMapEntry("PageTitle", html.EscapeString(fm.Title))
	hr.Cfg.Renderer.SetStringMapEntry("PageSubtitle", html.EscapeString(fm.Description))
	hr.Cfg.Renderer.SetDataMapEntry("FrontMatter", fm)
	hr.Cfg.Renderer.SetDataMapEntry("Content", hr.Cfg.MarkdownData["services.md"])
	err = hr.Cfg.Renderer.RenderTemplate(w, r, "services.gohtml")

	if err != nil {
		log.Fatalln(err)
	}
}
