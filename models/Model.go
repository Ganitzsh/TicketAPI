package models

import (
	"errors"
	"strconv"

	"github.com/SQLApi/mysql"
)

// User represents an actual person
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// GetAll returns all the users from the database
func (u *User) GetAll(selector []string) string {
	query := mysql.Select
	if selector != nil {
		count := len(selector)
		for i := 0; i < count; i++ {
			query += u.GetTableName() + "." + selector[i]
			if i != count-1 {
				query += mysql.Comma
			}
		}
	} else {
		query += mysql.WildcardAll
	}
	query += mysql.From + u.GetTableName()
	return query + mysql.EndQuery
}

// GetByID returns a user with specific ID
func (u *User) GetByID(id int, selector []string) string {
	sID := strconv.Itoa(id)
	query := mysql.Select
	if selector != nil {
		count := len(selector)
		for i := 0; i < count; i++ {
			query += u.GetTableName() + "." + selector[i]
			if i != count-1 {
				query += mysql.Comma
			}
		}
	} else {
		query += mysql.WildcardAll
	}
	query += mysql.From + u.GetTableName() + mysql.Where + u.GetPrimaryKeyField() + " = " + sID
	return query + mysql.EndQuery
}

// GetWhere returns one or many users corresponding to  wehre = values
func (u *User) GetWhere(selector, fileds, values []string) (string, error) {
	err := checkWhereConsistency(fileds, values)
	if err != nil {
		return "", err
	}
	query := mysql.Select
	if selector != nil {
		count := len(selector)
		for i := 0; i < count; i++ {
			query += u.GetTableName() + "." + selector[i]
			if i != count-1 {
				query += mysql.Comma
			}
		}
	} else {
		query += mysql.WildcardAll
	}
	query += mysql.From + u.GetTableName() + mysql.Where
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
func (u *User) GetTableName() string {
	return "user"
}

// GetPrimaryKeyField returns the primary key field name for this table
func (u *User) GetPrimaryKeyField() string {
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
func (u *User) GetInnerJoin(selector []string, joinedTable string, onRootColumn, onJoinedColumn string) string {
	query := mysql.Select
	if selector != nil {
		count := len(selector)
		for i := 0; i < count; i++ {
			query += u.GetTableName() + "." + selector[i]
			if i != count-1 {
				query += mysql.Comma
			}
		}
	} else {
		query += mysql.WildcardAll
	}
	query += mysql.From + u.GetTableName() + mysql.InnerJoin + joinedTable + mysql.On
	query += u.GetTableName() + "." + onRootColumn + mysql.Equals + joinedTable + "." + onJoinedColumn
	return query + mysql.EndQuery
}
