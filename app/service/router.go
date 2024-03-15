package service

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routing(e *echo.Echo, api *API) error {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	echoAPI := e.Group("/api")

	{
		v1 := echoAPI.Group("/v1")

		{
			pingAPI := v1.Group("/ping")

			pingAPI.GET("/", api.Ping)
		}
	}

	return nil
}
