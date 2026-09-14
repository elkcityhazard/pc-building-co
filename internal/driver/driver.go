package driver

import (
	"database/sql"

	"github.com/elkcityhazard/pc-building-company/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func NewDBConn(app *config.AppConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", app.DSN)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil

}
