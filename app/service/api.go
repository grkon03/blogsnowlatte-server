package service

import "github.com/grkon03/blogsnowlatte-server/implementation/database"

type API struct {
	Repo *database.Repository
}

func NewAPI(repo *database.Repository) API {
	return API{Repo: repo}
}
