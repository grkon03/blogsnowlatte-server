package database

type Repository struct {
	sqlh SQLHandler
}

func NewRepository(sqlh SQLHandler) *Repository {
	return &Repository{sqlh}
}
