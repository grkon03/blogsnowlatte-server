package infrastructure

import (
	"fmt"

	"github.com/grkon03/blogsnowlatte-server/implementation/database"
	"github.com/grkon03/blogsnowlatte-server/util/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SQLHandler struct {
	DB *gorm.DB
}

func NewSQLHandler(db *gorm.DB) *SQLHandler {
	return &SQLHandler{
		DB: db,
	}
}

func dsn(sc config.SQLConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		sc.User(),
		sc.Pass(),
		sc.Host(),
		sc.Port(),
		sc.Name(),
	)
}

func ConnectSQL(sc config.SQLConfig) (*SQLHandler, error) {
	db, err := gorm.Open(mysql.Open(dsn(sc)), &gorm.Config{})
	return NewSQLHandler(db), err
}

func (sqlh *SQLHandler) First(result interface{}, where ...interface{}) database.SQLHandler {
	db := sqlh.DB.First(result, where)
	return NewSQLHandler(db)
}

func (sqlh *SQLHandler) Create(value interface{}) database.SQLHandler {
	db := sqlh.DB.Create(value)
	return NewSQLHandler(db)
}

func (sqlh *SQLHandler) Where(query interface{}, args ...interface{}) database.SQLHandler {
	db := sqlh.DB.Where(query, args)
	return NewSQLHandler(db)
}

func (sqlh *SQLHandler) Error() error {
	return sqlh.DB.Error
}
