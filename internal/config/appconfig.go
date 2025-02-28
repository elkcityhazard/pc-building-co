package config

import (
	"database/sql"
	"html/template"
	"log"
	"sync"

	amrenderengine "github.com/elkcityhazard/am-render-engine"
	"github.com/elkcityhazard/pc-building-company/internal/models"
	"github.com/elkcityhazard/pc-building-company/pkg/mailer"
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
	MailMsgChan     chan *mailer.MailMessage
	MailErrChan     chan error
	MailDoneChan    chan bool
	WG              sync.WaitGroup
	SMTPHost        string
	SMTPPort        int
	SMTPUsername    string
	SMTPPassword    string
	SMTPFromAddr    string
	SMTPToAddr      string
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
		MailMsgChan:     make(chan *mailer.MailMessage),
		MailErrChan:     make(chan error),
		MailDoneChan:    make(chan bool),
		WG:              sync.WaitGroup{},
	}
}

func (a *AppConfig) ListenForErrors() {

	defer a.WG.Done()

	for {
		select {
		case err := <-a.MailErrChan:
			log.Println(err)
		}
	}
}
