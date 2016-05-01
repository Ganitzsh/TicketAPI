package config

import (
	"strconv"

	"github.com/SQLApi/mysql"
	"github.com/Sirupsen/logrus"
)

// SQLConfig is the representation of the SQL Server settings
type SQLConfig struct {
	Username string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	DBName   string `json:"database"`
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
}

// GetConnString returns the connection string to the database
func (c *SQLConfig) GetConnString() string {
	port := strconv.Itoa(c.Port)
	dsn := c.Username + ":" + c.Password + "@" + c.Protocol + "(" + c.Host + ":" + port + ")/" + c.DBName
	logrus.Info(dsn)
	return dsn
}

// NewConfig is
func NewConfig(user, pass, host, db, protocol string, port int) (*SQLConfig, error) {
	conf := SQLConfig{
		user,
		pass,
		host,
		db,
		protocol,
		port,
	}
	return &conf, nil
}

// GetDBType returns a string containing the name of the database type
func (c *SQLConfig) GetDBType() string {
	return mysql.Name
}
