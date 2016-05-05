package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"

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
func (h *MySQLHandler) GetAll(t interface{}, o ...interface{}) (interface{}, error) {
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
	rows, err := h.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	rowsCpy, err := h.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	// Use reflection on t object pased as parameter to automatically bind model's
	// struct field addresses for binding
	typeOf := reflect.TypeOf(t)
	var ptrs [][]interface{}
	var res []interface{}
	for rowsCpy.Next() {
		// Create new instance of object corresponding to model passed as parameter
		tmp := reflect.New(typeOf).Interface()
		fmt.Println("Instance type: ", reflect.TypeOf(tmp))
		// Add it to to result, to return the whole content of the request after
		// binding
		res = append(res, tmp)
		// Get reflect.Value of instace to loop through fields
		v := reflect.ValueOf(tmp).Elem()
		fieldAddresses := make([]interface{}, v.NumField())
		// Loop through instnce field
		for i := 0; i < v.NumField(); i++ {
			// Get current field reflect.Value
			valueField := v.Field(i)
			if valueField.CanAddr() {
				// Get the address of the given field
				tmp := valueField.Addr().Interface()
				fieldAddresses[i] = tmp
			}
		}
		// Store fields addresses into final array
		ptrs = append(ptrs, fieldAddresses)
	}
	// Bind model
	h.Bind(rows, ptrs)

	// DEBUG: Print result content
	for i := 0; i < len(res); i++ {
		fmt.Println("Value: ", res[i])
	}
	return res, nil
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
	tableName := o[0].(string)
	obj := o[1]
	query, args, err := builder.NewMySQLQuery(tableName, "").Insert(obj)
	if err != nil {
		return nil, err
	}
	stmt, err := h.Conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	return stmt.Exec(args...)
}

func (h *MySQLHandler) Update(o ...interface{}) (interface{}, error) {
	if o == nil {
		return nil, errors.New("No data to update")
	}
	tableName := o[0].(string)
	new := o[1]
	old := o[2]
	query, args, err := builder.NewMySQLQuery(tableName, "").Update(new, old)
	if err != nil {
		return nil, err
	}
	stmt, err := h.Conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	logrus.WithField("args", args).Info(query)
	return stmt.Exec(args...)
}

func (h *MySQLHandler) Delete(o ...interface{}) (interface{}, error) {
	if o == nil {
		return nil, errors.New("No data to delete")
	}
	tableName := o[0].(string)
	obj := o[1]
	query, args, err := builder.NewMySQLQuery(tableName, "").Delete(obj)
	if err != nil {
		return nil, err
	}
	stmt, err := h.Conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	logrus.WithField("args", args).Info(query)
	return stmt.Exec(args...)
}

func (h *MySQLHandler) Bind(content *sql.Rows, ptrs [][]interface{}) {
	i := 0
	for content.Next() {
		err := content.Scan(ptrs[i]...)
		if err != nil {
			panic(err)
		}
		fmt.Println(ptrs[i])
		i++
	}
}
