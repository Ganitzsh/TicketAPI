package config

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
)

// APIConfig is
type APIConfig struct {
	ListeningPort int             `json:"port"`
	DBType        string          `json:"db_type"`
	DBSettings    json.RawMessage `json:"db_settings"`
}

// NewAPIConfig is
func NewAPIConfig(port int, dbType string) (*APIConfig, error) {
	conf := APIConfig{
		ListeningPort: port,
		DBType:        dbType,
	}
	return &conf, nil
}

// NewAPIConfigFromFile returs
func NewAPIConfigFromFile(path string) (*APIConfig, error) {
	var conf APIConfig
	file, e := ioutil.ReadFile(path)
	if e != nil {
		log.WithFields(log.Fields{
			"error": e.Error(),
			"path":  path,
		}).Error("Error while opening config file")
		return nil, e
	}
	json.Unmarshal(file, &conf)
	return &conf, nil
}
