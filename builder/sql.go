package builder

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/SQLApi/mysql"
)

// MySQLQueryBuilder is the main query builder for MySQL flavored DB
type MySQLQueryBuilder struct {
	tableName  string
	primaryKey string
}

// NewMySQLQuery returns a new instance of MySQLQueryBuilder
func NewMySQLQuery(tableName, primaryKey string) *MySQLQueryBuilder {
	return &MySQLQueryBuilder{tableName, primaryKey}
}

// GetAll returns all the users from the database
func (qb *MySQLQueryBuilder) GetAll(selector []string) string {
	query := mysql.Select
	if selector != nil {
		count := len(selector)
		for i := 0; i < count; i++ {
			query += qb.tableName + "." + selector[i]
			if i != count-1 {
				query += mysql.Comma
			}
		}
	} else {
		query += mysql.WildcardAll
	}
	query += mysql.From + qb.tableName
	return query + mysql.EndQuery
}

// GetByID returns a user with specific ID
func (qb *MySQLQueryBuilder) GetByID(id int, selector []string) string {
	sID := strconv.Itoa(id)
	query := mysql.Select
	if selector != nil {
		count := len(selector)
		for i := 0; i < count; i++ {
			query += qb.tableName + "." + selector[i]
			if i != count-1 {
				query += mysql.Comma
			}
		}
	} else {
		query += mysql.WildcardAll
	}
	query += mysql.From + qb.tableName + mysql.Where + qb.primaryKey + mysql.Equals + sID
	return query + mysql.EndQuery
}

// GetWhere returns one or many users corresponding to  wehre = values
func (qb *MySQLQueryBuilder) GetWhere(selector, fileds, values []string) (string, error) {
	err := checkWhereConsistency(fileds, values)
	if err != nil {
		return "", err
	}
	query := mysql.Select
	if selector != nil {
		count := len(selector)
		for i := 0; i < count; i++ {
			query += qb.tableName + "." + selector[i]
			if i != count-1 {
				query += mysql.Comma
			}
		}
	} else {
		query += mysql.WildcardAll
	}
	query += mysql.From + qb.tableName + mysql.Where
	count := len(fileds)
	for j := 0; j < count; j++ {
		query += fileds[j] + mysql.Equals + "\"" + values[j] + "\""
		if j != count-1 {
			query += mysql.And
		}
	}
	return query + mysql.EndQuery, nil
}

// GetTableName returns the name of the users table
func (qb *MySQLQueryBuilder) GetTableName() string {
	return "user"
}

// GetPrimaryKeyField returns the primary key field name for this table
func (qb *MySQLQueryBuilder) GetPrimaryKeyField() string {
	return "id"
}

// checkWhereConsistency checks if the where and values arrays are consistent
// with each other
func checkWhereConsistency(fields, values []string) error {
	if fields == nil {
		return errors.New("Desired field not specified")
	}
	if values == nil {
		return errors.New("Desired values not specified")
	}
	lenFields := len(fields)
	lenValues := len(values)
	if lenFields != lenValues {
		return errors.New("Inconsistent number of values in arrays")
	}
	return nil
}

// GetInnerJoin returns
func (qb *MySQLQueryBuilder) GetInnerJoin(selector []string, joinedTable string, onRootColumn, onJoinedColumn string) string {
	query := mysql.Select
	if selector != nil {
		count := len(selector)
		for i := 0; i < count; i++ {
			query += qb.tableName + "." + selector[i]
			if i != count-1 {
				query += mysql.Comma
			}
		}
	} else {
		query += mysql.WildcardAll
	}
	query += mysql.From + qb.tableName + mysql.InnerJoin + joinedTable + mysql.On
	query += qb.tableName + "." + onRootColumn + mysql.Equals + joinedTable + "." + onJoinedColumn
	return query + mysql.EndQuery
}

// Insert returns
func (qb *MySQLQueryBuilder) Insert(o interface{}) (string, error) {
	query := mysql.InsertInto + qb.tableName + mysql.Values + "("
	v := reflect.ValueOf(o)
	for i := 0; i < v.NumField(); i++ {
		t := v.Field(i).Interface()
		switch t := t.(type) {
		case string:
			query += "\"" + t + "\""
		case int:
			query += strconv.Itoa(t)
		case float64:
			number := strconv.FormatFloat(t, 'e', -1, 64)
			s := strings.Split(number, "e")
			if len(s) < 2 {
				query += number
			} else {
				query += s[0]
			}
		}
		if i != v.NumField()-1 {
			query += mysql.Comma
		}
	}
	query += ")"
	return query + mysql.EndQuery, nil
}

// Delete returns
func (qb *MySQLQueryBuilder) Delete(o interface{}) (string, error) {
	query := mysql.Delete + mysql.From + qb.tableName + mysql.Where
	v := reflect.ValueOf(o)
	w := reflect.Indirect(v)
	st := reflect.TypeOf(o)
	for i := 0; i < v.NumField(); i++ {
		field := st.Field(i)
		name := field.Tag.Get("db")
		if name == "" {
			query += w.Type().Field(i).Name + mysql.Equals
		} else {
			query += name + mysql.Equals
		}
		t := v.Field(i).Interface()
		switch t := t.(type) {
		case string:
			query += "\"" + t + "\""
		case int:
			query += strconv.Itoa(t)
		case float64:
			number := strconv.FormatFloat(t, 'e', -1, 64)
			s := strings.Split(number, "e")
			if len(s) < 2 {
				query += number
			} else {
				query += s[0]
			}
		}
		if i != v.NumField()-1 {
			query += mysql.And
		}
	}
	return query + mysql.EndQuery, nil
}

func (qb *MySQLQueryBuilder) Update(new, old interface{}) (string, error) {
	query := mysql.Update + qb.tableName + mysql.Set
	v := reflect.ValueOf(new)
	w := reflect.Indirect(v)
	st := reflect.TypeOf(new)
	for i := 0; i < v.NumField(); i++ {
		field := st.Field(i)
		name := field.Tag.Get("db")
		if name == "" {
			query += w.Type().Field(i).Name + mysql.Equals
		} else {
			query += name + mysql.Equals
		}
		t := v.Field(i).Interface()
		switch t := t.(type) {
		case string:
			query += "\"" + t + "\""
		case int:
			query += strconv.Itoa(t)
		case float64:
			number := strconv.FormatFloat(t, 'e', -1, 64)
			s := strings.Split(number, "e")
			if len(s) < 2 {
				query += number
			} else {
				query += s[0]
			}
		}
		if i != v.NumField()-1 {
			query += mysql.Comma
		}
	}
	query += mysql.Where
	v = reflect.ValueOf(old)
	w = reflect.Indirect(v)
	st = reflect.TypeOf(old)
	for i := 0; i < v.NumField(); i++ {
		field := st.Field(i)
		name := field.Tag.Get("db")
		if name == "" {
			query += w.Type().Field(i).Name + mysql.Equals
		} else {
			query += name + mysql.Equals
		}
		t := v.Field(i).Interface()
		switch t := t.(type) {
		case string:
			query += "\"" + t + "\""
		case int:
			query += strconv.Itoa(t)
		case float64:
			number := strconv.FormatFloat(t, 'e', -1, 64)
			s := strings.Split(number, "e")
			if len(s) < 2 {
				query += number
			} else {
				query += s[0]
			}
		}
		if i != v.NumField()-1 {
			query += mysql.Comma
		}
	}
	return query + mysql.EndQuery, nil
}
