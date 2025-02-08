package main

import (
	"flag"

	"github.com/elkcityhazard/pc-building-company/internal/config"
)

func parseFlags(app *config.AppConfig) {

	flag.StringVar(&app.WebsiteName, "website_name", "My Cool Website", "the name you want to give to your website")
	flag.StringVar(&app.WebsiteAddress, "website_url", "http://localhost", "the base url for the site")
	flag.BoolVar(&app.IsProduction, "is_production", false, "set environment(development|production)")
	flag.Parse()

}
