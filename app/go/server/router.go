package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter() (*echo.Echo, error) {

	router := echo.New()

	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "minimal_sns_app")
	})

	return router, nil
}
