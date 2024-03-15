package database

type SQLHandler interface {
	First(result interface{}, where ...interface{}) SQLHandler
	Create(value interface{}) SQLHandler
	Where(query interface{}, args ...interface{}) SQLHandler
	Error() error
}
