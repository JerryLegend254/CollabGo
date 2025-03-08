package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *application) pingHandler(c echo.Context) error {
	return app.jsonResponse(c, http.StatusOK, map[string]string{
		"status": "ok",
	})
}
