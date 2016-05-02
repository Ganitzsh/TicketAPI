package handlers

// IHandler is the main interface for handlers
type IHandler interface {
	IsAvailable() error
	GetAll(o ...interface{}) (interface{}, error)
	GetBy(o ...interface{}) (interface{}, error)
	Insert(o ...interface{}) (interface{}, error)
	Update(o ...interface{}) (interface{}, error)
	Delete(o ...interface{}) (interface{}, error)
}
