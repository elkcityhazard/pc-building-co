package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	amrenderengine "github.com/elkcityhazard/am-render-engine"
	"github.com/elkcityhazard/pc-building-company/content"
	"github.com/elkcityhazard/pc-building-company/internal/config"
	"github.com/elkcityhazard/pc-building-company/internal/handlers"
	"github.com/elkcityhazard/pc-building-company/internal/models"
	"github.com/elkcityhazard/pc-building-company/internal/templates"
	"github.com/elkcityhazard/pc-building-company/pkg/mailer"
	"github.com/yuin/goldmark"
)

var app *config.AppConfig = config.NewAppConfig()

func main() {
	app.Renderer = amrenderengine.NewTemplateCollection(templates.GetTemplatesFS(), "./internal/templates")
	app.Renderer.CreateTemplateCache()
	app.Renderer.SetDirPaths("templates/pages", "templates/layouts", "templates/partials", "gohtml")
	app.Renderer.TemplateFuncs = template.FuncMap{
		"humanPhone": func(ph string) string {
			out := ph[:2]
			out += "("
			out += ph[2:5]
			out += ")"
			out += " "
			out += ph[5:8]
			out += "-"
			out += ph[8:12]
			return html.EscapeString(out)

		},
		"safeHTML": func(markup string) template.HTML {
			return template.HTML(markup)
		},
		"getYear": func() int {
			return time.Now().Year()
		},
		"concatBaseURL": func(s string) string {
			return app.WebsiteAddress + s
		},
	}

	_, err := app.Renderer.CreateTemplateCache()

	if err != nil {
		log.Fatalln(err)
	}

	parseFlags(app)

	app.Renderer.UseCache = app.IsProduction
	app.Renderer.SetIsProduction(app.IsProduction)
	app.Renderer.SetStringMapEntry("SiteTitle", app.WebsiteName)
	if app.IsProduction {
		app.Renderer.SetStringMapEntry("BaseURL", app.WebsiteAddress)
	} else {
		app.Renderer.SetStringMapEntry("BaseURL", fmt.Sprintf("%s%s", app.WebsiteAddress, app.Port))
	}

	serviceFile, err := os.Open(filepath.Join("./static/data", "services.json"))

	if err != nil {
		log.Fatalln(err)
	}

	defer serviceFile.Close()

	var services models.HomeServices

	sData, err := io.ReadAll(serviceFile)

	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(sData, &services)

	if err != nil {
		log.Fatalln(err)
	}

	app.Renderer.SetStringMapEntry("ContactLink", "/contact#contactForm")
	app.Renderer.SetStringMapEntry("PhoneNumber", "+12313576340")
	app.Renderer.SetDataMapEntry("Services", services)

	parseMarkdownSet(app)

	newMainNavigation(app)
	err = generateServiceJSON(app)

	if err != nil {
		log.Fatalln(err)
	}

	handlerRepo := handlers.NewHandlerRepo(app, nil)
	handlers.SetHandlerRepo(handlerRepo)

	mailer := mailer.NewMailer(app.SMTPHost, app.SMTPUsername, app.SMTPPassword, app.SMTPPort)

	mailer.NewDialer()

	app.WG.Add(1)
	go mailer.ListenForMail(app.MailMsgChan, app.MailErrChan, app.MailDoneChan, &app.WG)

	app.WG.Add(1)
	go app.ListenForErrors()

	srv := &http.Server{
		Addr:              app.Port,
		Handler:           routes(),
		IdleTimeout:       time.Second * 30,
		ReadTimeout:       time.Second * 30,
		ReadHeaderTimeout: time.Second * 30,
		WriteTimeout:      time.Second * 30,
	}

	log.Println("Starting Server on", app.Port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}

}

func newMainNavigation(app *config.AppConfig) {
	navList := models.NewNavList()

	navList.New(models.NewNavItem("Services", "/services", 0))
	navList.NavItems = append(navList.NavItems, models.NewNavItem("Contact", "/contact", 1))
	app.Renderer.SetDataMapEntry("MainNav", navList)
}

func generateServiceJSON(app *config.AppConfig) error {

	var (
		pathToServices string = "./static/data/services.json"
	)
	f, err := os.Open(pathToServices)

	if err != nil {
		return err
	}

	defer f.Close()

	fData, err := io.ReadAll(f)

	if err != nil {
		return err
	}
	var services models.HomeServices

	err = json.Unmarshal(fData, &services)

	if err != nil {
		return err
	}

	app.Services = services
	app.Renderer.SetDataMapEntry("Services", services)

	return nil

}

func parseMarkdownSet(app *config.AppConfig) {
	mdSet := make(map[string]template.HTML)
	contentFS := content.GetContentFS()
	baseDir, err := contentFS.ReadDir("content")
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range baseDir {

		name := v.Name()

		buf := bytes.Buffer{}

		file, err := contentFS.ReadFile(filepath.Join("content", name))

		if err != nil {
			log.Fatalln(err)
		}

		if strings.LastIndex(string(file), "---\n") == -1 {
			continue
		}

		body := string(file)[strings.LastIndex(string(file), "---\n")+4:]

		err = goldmark.Convert([]byte(body), &buf)

		if err != nil {
			log.Fatalln(err)
		}

		mdSet[name] = template.HTML(buf.String())

	}

	app.MarkdownData = mdSet

}
