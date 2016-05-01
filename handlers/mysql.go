package handlers

import (
	"database/sql"
	"errors"

	"github.com/SQLApi/builder"
	"github.com/SQLApi/config"
	"github.com/SQLApi/mysql"
	"github.com/Sirupsen/logrus"
)

// MySQLHandler is the main handler for the MySQL database type
type MySQLHandler struct {
	Conn   *sql.DB
	DBName string
}

// NewMySQLHandler returns a pointer to a new MySQLHandler
func NewMySQLHandler(c *config.MySQLConfig) (IHandler, error) {
	db, err := sql.Open(mysql.Name, c.GetConnString())
	if err != nil {
		return nil, err
	}
	h := MySQLHandler{
		Conn: db,
	}
	return &h, nil
}

// IsAvailable checks if the database is still available
func (h *MySQLHandler) IsAvailable() error {
	return h.Conn.Ping()
}

// GetAll retrieves all elements of the given table from the database
func (h *MySQLHandler) GetAll(o ...interface{}) (interface{}, error) {
	var selector []string
	if o == nil {
		return nil, errors.New("No table name specified")
	}
	tableName := o[0].(string)
	if tableName == "" {
		return nil, errors.New("No table name specified")
	}
	if len(o) > 1 {
		selector = o[1].([]string)
	}
	query := builder.NewMySQLQuery(tableName, "").GetAll(selector)
	logrus.WithField("query", query).Info("Querying " + tableName + " table")
	return h.Conn.Query(query)
}

func (h *MySQLHandler) GetBy(o ...interface{}) (interface{}, error) {
	var selector []string
	if o == nil {
		return nil, errors.New("No filter specified")
	}
	if len(o) < 4 {
		return nil, errors.New("Not enough arguments specified for query")
	}
	tableName := o[0].(string)
	if o[1] != nil {
		selector = o[1].([]string)
	}
	fields := o[2].([]string)
	values := o[3].([]string)
	query, err := builder.NewMySQLQuery(tableName, "").GetWhere(selector, fields, values)
	if err != nil {
		return nil, err
	}
	logrus.WithField("query", query).Info("Querying " + tableName + " table")
	return h.Conn.Query(query)
}

func (h *MySQLHandler) Insert(o ...interface{}) (interface{}, error) {
	if o == nil {
		return nil, errors.New("No data to insert")
	}
	return nil, nil
}

func (h *MySQLHandler) Update(o ...interface{}) (interface{}, error) {
	return nil, nil
}

func (h *MySQLHandler) Delete(o ...interface{}) {

}
