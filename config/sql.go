package config

// SQLConfig is the representation of the SQL Server settings
type SQLConfig struct {
	Username string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	DBName   string `json:"database"`
}

func (c *SQLConfig) getConnString() string {
	return c.Username + ":" + c.Password + "@" + c.Host + "/" + c.DBName
}

// NewConfig is
func NewConfig(user, pass, host, db string) {

}

// NewConfigFromFile instantiate a new SQLConfig using a give file
func NewConfigFromFile(path string) (*SQLConfig, error) {
	return nil, nil
}
