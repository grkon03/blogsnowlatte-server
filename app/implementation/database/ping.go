package database

import (
	"github.com/grkon03/blogsnowlatte-server/implementation/dataset"
	"github.com/grkon03/blogsnowlatte-server/model"
)

// No args; Ping table is expected only a record
func (repo *Repository) GetPing() (*dataset.Ping, error) {
	var result model.Ping
	err := repo.sqlh.Where(&model.Ping{ID: "ping"}).First(&result).Error()

	if err != nil {
		return nil, convertErr(err)
	}

	ping := &dataset.Ping{
		ID:    result.ID,
		Value: result.Value,
	}

	return ping, nil
}

// No args; Ping table is expected only a record
func (repo *Repository) CreatePing() error {
	err := repo.sqlh.Create(&model.Ping{
		ID:    "ping",
		Value: "pong",
	}).Error()

	if err != nil {
		return convertErr(err)
	}

	return nil
}
