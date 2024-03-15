package database

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("record not found")
)

func convertErr(err error) error {
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}

	return err
}
