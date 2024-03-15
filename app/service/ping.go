package service

import (
	"net/http"

	"github.com/grkon03/blogsnowlatte-server/implementation/database"
	"github.com/labstack/echo/v4"
)

// GET /ping
func (api *API) Ping(c echo.Context) error {
	ping, err := api.Repo.GetPing()
	if err == database.ErrNotFound {
		api.Repo.CreatePing()
		ping, err = api.Repo.GetPing()
		if err != nil {
			return c.String(http.StatusNotImplemented, "ping missed")
		}
	}

	return c.String(http.StatusOK, ping.Value)
}
