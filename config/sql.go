package config

import (
	"strconv"

	"github.com/SQLApi/mysql"
	"github.com/Sirupsen/logrus"
)

// MySQLConfig is the representation of the SQL Server settings
type MySQLConfig struct {
	Username string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	DBName   string `json:"database"`
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
}

// GetConnString returns the connection string to the database
func (c *MySQLConfig) GetConnString() string {
	port := strconv.Itoa(c.Port)
	dsn := c.Username + ":" + c.Password + "@" + c.Protocol + "(" + c.Host + ":" + port + ")/" + c.DBName
	logrus.Info(dsn)
	return dsn
}

// NewMySQLConfig is
func NewMySQLConfig(user, pass, host, db, protocol string, port int) (*MySQLConfig, error) {
	conf := MySQLConfig{
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
func (c *MySQLConfig) GetDBType() string {
	return mysql.Name
}
