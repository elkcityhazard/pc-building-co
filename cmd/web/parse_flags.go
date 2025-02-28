package main

import (
	"flag"

	"github.com/elkcityhazard/pc-building-company/internal/config"
)

func parseFlags(app *config.AppConfig) {

	flag.StringVar(&app.WebsiteName, "website_name", "My Cool Website", "the name you want to give to your website")
	flag.StringVar(&app.WebsiteAddress, "website_url", "http://localhost", "the base url for the site")
	flag.BoolVar(&app.IsProduction, "is_production", false, "set environment(development|production)")
	flag.StringVar(&app.Port, "port", ":8080", "the port the application runs on (:6587)")

	flag.StringVar(&app.SMTPHost, "smtp_host", "localhost", "the hostname for the mail server")
	flag.IntVar(&app.SMTPPort, "smtp_port", 1025, "the mail server port")
	flag.StringVar(&app.SMTPUsername, "smtp_username", "username", "the username credential for the smtp server")
	flag.StringVar(&app.SMTPPassword, "smtp_password", "password", "the password credential for the smtp server")
	flag.StringVar(&app.SMTPToAddr, "smtp_toaddr", "test@test.com", "the address that mail from the server will be sent to")
	flag.StringVar(&app.SMTPFromAddr, "smtp_fromaddr", "web@testing.com", "the address that will be put into the from field of the email")

	flag.Parse()

}
