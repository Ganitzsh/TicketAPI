package config

// APIConfig is
type APIConfig struct {
	ListeningPort int        `json:"port"`
	DBSettings    *SQLConfig `json:"db_settings"`
}

// NewAPIConfig is
func NewAPIConfig(port int, dbSettings *SQLConfig) (*APIConfig, error) {
	return nil, nil
}

// NewAPIConfigFromFile returs
func NewAPIConfigFromFile(path string) (*APIConfig, error) {
	return nil, nil
}
