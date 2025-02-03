package handlers

import (
	"github.com/elkcityhazard/pc-building-company/internal/config"
	"github.com/elkcityhazard/pc-building-company/internal/repository"
)

var (
	Repo             *HandlerRepo
	pathToServices   string = "./static/data/services.json"
	markdownBasePath string = "./content"
)

type HandlerRepo struct {
	Cfg  *config.AppConfig
	Conn repository.DbServicer
}

func NewHandlerRepo(a *config.AppConfig, c repository.DbServicer) *HandlerRepo {
	return &HandlerRepo{
		Cfg:  a,
		Conn: c,
	}
}

func SetHandlerRepo(hr *HandlerRepo) {
	Repo = hr
}
