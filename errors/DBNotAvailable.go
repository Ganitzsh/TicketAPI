package errors

// DBNotAvailable is ...
type DBNotAvailable struct {
	s string
}

// NewDBNotAvailable instantiate a new error
func NewDBNotAvailable(text string) error {
	return &DBNotAvailable{text}
}

func (e *DBNotAvailable) Error() string {
	return e.s
}
