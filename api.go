package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/SQLApi/config"
	"github.com/SQLApi/controllers"
	"github.com/SQLApi/errors"
	"github.com/SQLApi/handlers"
	log "github.com/Sirupsen/logrus"
)

// API is the main type 'object' containing every element to work properly
type API struct {
	DBHandler handlers.IHandler
	Settings  *config.APIConfig
	Router    *gin.Engine
}

// NewAPI is the constructor of the API
func NewAPI(port int, dbType string, dbs map[string]interface{}) (*API, error) {
	var handler handlers.IHandler
	settings, err := config.NewAPIConfig(port, dbType)
	if err != nil {
		return nil, err
	}
	switch settings.DBType {
	case "":
		return nil, errors.NewError("No database type specified")
	case "mysql":
		tmp := config.MySQLConfig{
			Username: dbs["user"].(string),
			Password: dbs["password"].(string),
			Host:     dbs["host"].(string),
			DBName:   dbs["database"].(string),
			Protocol: dbs["protocol"].(string),
			Port:     dbs["port"].(int),
		}
		handler, err = handlers.NewMySQLHandler(&tmp)
	default:
		return nil, errors.NewUnsupportedDatabase("Unsuported database type " + settings.DBType)
	}
	if err != nil {
		log.Error("Could not get a handle to the database")
		return nil, err
	}
	api := API{
		DBHandler: handler,
		Settings:  settings,
		Router:    gin.Default(),
	}
	api.configureRoutes()
	if !api.available() {
		return nil, errors.NewDBNotAvailable("Database not available")
	}
	return &api, nil
}

// NewAPIFromFile is the constructor of the API
func NewAPIFromFile(path string) (*API, error) {
	var handler handlers.IHandler
	settings, err := config.NewAPIConfigFromFile(path)
	if err != nil {
		return nil, err
	}
	fmt.Println(settings)
	switch settings.DBType {
	case "":
		return nil, errors.NewError("No database type specified")
	case "mysql":
		tmp := new(config.MySQLConfig)
		if err = json.Unmarshal(settings.DBSettings, tmp); err != nil {
			panic(err)
		}
		handler, err = handlers.NewMySQLHandler(tmp)
	default:
		return nil, errors.NewUnsupportedDatabase("Unsuported database type " + settings.DBType)
	}
	if err != nil {
		log.Error("Could not get a handle to the database")
		return nil, err
	}
	api := API{
		DBHandler: handler,
		Settings:  settings,
		Router:    gin.Default(),
	}
	api.configureRoutes()
	if !api.available() {
		return nil, errors.NewDBNotAvailable("Database not available")
	}
	return &api, nil
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
	err := api.DBHandler.IsAvailable()
	if err != nil {
		log.WithField("error", err.Error()).Error("Database not available")
		return false
	}
	return true
}
