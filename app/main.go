package main

import (
	"log"

	"github.com/grkon03/blogsnowlatte-server/implementation/database"
	"github.com/grkon03/blogsnowlatte-server/infrastructure"
	"github.com/grkon03/blogsnowlatte-server/service"
	"github.com/grkon03/blogsnowlatte-server/util/config"
	"github.com/labstack/echo/v4"
)

func main() {

	appConf, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	var sqlh *infrastructure.SQLHandler
	sqlh, err = infrastructure.ConnectSQL(appConf.SQLConfig())
	if err != nil {
		log.Fatal(err)
	}

	if appConf.ExecConfig().DoMiguration() {
		sqlh.AutoMigration()
	}

	repo := database.NewRepository(sqlh)
	api := service.NewAPI(repo)

	e := echo.New()

	err = service.Routing(e, &api)

	if err != nil {
		log.Fatal(err)
	}

	e.Logger.Fatal(e.Start(":3111"))
}
