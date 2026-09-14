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
	"github.com/elkcityhazard/pc-building-company/internal/driver"
	"github.com/elkcityhazard/pc-building-company/internal/handlers"
	"github.com/elkcityhazard/pc-building-company/internal/models"
	"github.com/elkcityhazard/pc-building-company/internal/templates"
	"github.com/elkcityhazard/pc-building-company/pkg/mailer"
	"github.com/yuin/goldmark"
	"golang.org/x/image/webp"
)

var app *config.AppConfig = config.NewAppConfig()

type image_t struct {
	Name   string
	Alt    string
	Height int
	Width  int
}

func main() {
	app.Renderer = amrenderengine.NewTemplateCollection(templates.GetTemplatesFS(), "./internal/templates")
	app.Renderer.CreateTemplateCache()
	app.Renderer.UseCache = app.IsProduction
	app.Renderer.SetIsProduction(app.IsProduction)
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
		"getImages": func(s string) []image_t {
			var imgMeta []string = []string{
	"Custom closet remodel with built-in storage in Leelanau County, Michigan",
	"Custom built-in shelving installation for a home remodel in Leelanau County, Michigan",
	"Custom shelving and fireplace mantel built in Suttons Bay, Michigan",
	"Custom built-in storage bench designed for a home remodel in Lake Leelanau, Michigan",
	"Detailed view of custom built-in storage from a Lake Leelanau home remodeling project",
	"Custom built-in bookcase handcrafted in Leelanau County, Michigan",
	"Full view of custom built-in bookshelves made locally in Leelanau County",
	"Custom bathtub surround with detailed trim from a Leelanau County bathroom remodel",
	"Custom bathroom vanity, tilework, and cabinetry completed in Leelanau County, Michigan",
	"Walk-in shower with custom tilework and built-in bench from a Lake Leelanau bathroom remodel",
	"Custom tile walk-in shower and bathtub from a Leelanau County bathroom remodeling project",
	"Custom bathroom vanity, shelving, and tilework completed in Leelanau County, Michigan",
	"Custom outdoor deck remodel creating an inviting gathering space in Leelanau County, Michigan",
	"Custom deck remodel with new windows, trim, and painting in Lake Leelanau, Michigan",
	"Early construction phase of a living room remodel in Lake Leelanau, Michigan",
	"Living room remodel in progress at a Lake Leelanau, Michigan home",
	"Initial construction phase of a custom staircase project in Leelanau County, Michigan",
	"Completed custom staircase with stained wood, trim, and safety railing in Leelanau County",
	"Custom laundry room remodel with functional storage by P.C. Building Company in Leelanau County, Michigan",
	"Early stage of an exterior home restoration project in Leelanau County, Michigan",
	"Restored exterior trim, siding, and window inlay from a Lake Leelanau home remodel",
}

			var imgLst []image_t
			files, err := filepath.Glob(fmt.Sprintf("%s", s))
			if err != nil {
				fmt.Println(err)
				return nil
			}

			for j, v := range files {
				f, err := os.Open(v)
				if err != nil {
					continue
				}
				defer f.Close()

				img, err := webp.DecodeConfig(f)
				if err != nil {
					continue
				}

				var i image_t
				i.Name = fmt.Sprintf("%s", f.Name())
				i.Alt = imgMeta[j]
				i.Width = img.Width
				i.Height = img.Height

				imgLst = append(imgLst, i)
			}

			return imgLst
		},
	}

	_, err := app.Renderer.CreateTemplateCache()

	if err != nil {
		log.Fatalln(err)
	}

	parseFlags(app)

	db, err := driver.NewDBConn(app)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

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

	handlerRepo := handlers.NewHandlerRepo(app, db)
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
