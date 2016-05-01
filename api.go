package main

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/SQLApi/config"
	"github.com/SQLApi/controllers"
	"github.com/SQLApi/errors"
	log "github.com/Sirupsen/logrus"
)

// MySQLHandler is an handler for MySQL database
type MySQLHandler struct {
	Handler *sql.DB
}

// API is the main type 'object' containing every element to work properly
type API struct {
	DBHandler *sql.DB
	Settings  *config.APIConfig
	Router    *gin.Engine
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

// NewAPIFromFile is the constructor of the API
func NewAPIFromFile(path string) (*API, error) {
	var db *sql.DB
	settings, err := config.NewAPIConfigFromFile(path)
	if err != nil {
		return nil, err
	}
	switch settings.DBType {
	case "":
		return nil, errors.NewError("No database type specified")
	case "mysql":
		db, err = sql.Open(settings.DBType, settings.DBSettings.GetConnString())
	default:
		return nil, errors.NewUnsupportedDatabase("Unsuported database type " + settings.DBType)
	}
	if err != nil {
		log.Error("Could not get a handle to the database")
		return nil, err
	}
	api := &API{
		DBHandler: db,
		Settings:  settings,
		Router:    gin.Default(),
	}
	api.configureRoutes()
	if !api.available() {
		return nil, errors.NewDBNotAvailable("Database not available")
	}
	return api, nil
}

func (api *API) configureRoutes() {
	api.Router.GET("/tickets", controllers.Tickets)
	api.Router.POST("/ticket/:id", controllers.NewTicket)
	api.Router.DELETE("/ticket/:id", controllers.DeleteTicket)
	api.Router.PATCH("/ticket/:id", controllers.EditTicket)
}

// Run starts listening on given port
func (api *API) Run() {
	port := ":" + strconv.Itoa(api.Settings.ListeningPort)
	api.Router.Run(port)
}

//TODO: Add dbtype switch to check availability for multiple DB types
func (api *API) available() bool {
	err := api.DBHandler.Ping()
	if err != nil {
		log.WithField("error", err.Error()).Error("Database not available")
		return false
	}
	return true
}
