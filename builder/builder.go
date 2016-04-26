package builder

// IQueryBuilder is the interface type for every sub builders
type IQueryBuilder interface {
	GetAll(selector []string) string
	GetByID(id int, selector []string) string
	GetTableName() string
}
