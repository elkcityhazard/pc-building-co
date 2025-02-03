package config

import (
	"database/sql"
	"html/template"

	amrenderengine "github.com/elkcityhazard/am-render-engine"
	"github.com/elkcityhazard/pc-building-company/internal/models"
)

type AppConfig struct {
	WebsiteName     string
	WebsiteAddress  string
	Port            string
	Conn            *sql.DB
	Renderer        *amrenderengine.TemplateCollection
	TemplateData    map[string]*template.Template
	MarkdownData    map[string]template.HTML
	IsProduction    bool
	IsAuthenticated bool
	SiteTitle       string
	MsgChan         chan string
	ErrorChan       chan error
	DoneChan        chan bool
	Services        models.HomeServices
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		WebsiteName:     "",
		WebsiteAddress:  "",
		Port:            ":8080",
		Conn:            nil,
		Renderer:        nil,
		TemplateData:    nil,
		MarkdownData:    nil,
		IsProduction:    false,
		IsAuthenticated: false,
		SiteTitle:       "",
		MsgChan:         make(chan string),
		ErrorChan:       make(chan error),
		DoneChan:        make(chan bool),
		Services:        models.HomeServices{},
	}
}
