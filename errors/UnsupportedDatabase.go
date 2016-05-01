package errors

// NewUnsupportedDatabase is ...
type UnsupportedDatabase struct {
	s string
}

// NewUnsupportedDatabase instantiate a new error
func NewUnsupportedDatabase(text string) error {
	return &DBNotAvailable{text}
}

func (e *UnsupportedDatabase) Error() string {
	return e.s
}
