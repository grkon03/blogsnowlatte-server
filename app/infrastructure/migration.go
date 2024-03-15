package infrastructure

import "github.com/grkon03/blogsnowlatte-server/model"

func (sqlh *SQLHandler) AutoMigration() {
	autoMigrationV1(sqlh)
}

func autoMigrationV1(sqlh *SQLHandler) {
	sqlh.DB.AutoMigrate(
		&model.Ping{},
	)
}
