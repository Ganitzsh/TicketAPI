package main

import (
	"database/sql"

	"github.com/SQLApi/errors"
	log "github.com/Sirupsen/logrus"
)

// API is the main type 'object' containing every element to work properly
type API struct {
	DBHandler *sql.DB
}

// NewAPI is the constructor of the API
func NewAPI(port int32) (*API, error) {
	db, err := sql.Open("mysql", "")
	if err != nil {
		log.Error("Could not get a handle to the database")
		panic(err)
	}
	api := &API{
		DBHandler: db,
	}
	if !api.available() {
		return nil, errors.NewDBNotAvailable("Database not available")
	}
	return api, nil
}

func (api *API) available() bool {
	err := api.DBHandler.Ping()
	if err != nil {
		log.WithField("error", err.Error()).Error("Database not available")
		return false
	}
	return true
}
