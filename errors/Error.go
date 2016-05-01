package errors

// Error is ...
type Error struct {
	s string
}

// NewError instantiate a new error
func NewError(text string) error {
	return &Error{text}
}

func (e *Error) Error() string {
	return e.s
}
