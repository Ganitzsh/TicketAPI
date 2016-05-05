package handlers

// IHandler is the main interface for handlers
type IHandler interface {
	IsAvailable() error
	GetAll(t interface{}, o ...interface{}) (interface{}, error)
	GetBy(t interface{}, o ...interface{}) (interface{}, error)
	Insert(o ...interface{}) (interface{}, error)
	Update(o ...interface{}) (interface{}, error)
	Delete(o ...interface{}) (interface{}, error)
}
