package config

// IDBConfig is an interface representing a database configuration
type IDBConfig interface {
	GetDBType() string
}
