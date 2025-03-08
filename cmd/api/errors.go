package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *application) internalServerError(c echo.Context, err error) error {
	app.logger.Errorw("internal server error", "error", err.Error(), "method", c.Request().Method, "path", c.Request().URL.Path)
	return writeJSONError(c, http.StatusInternalServerError, "something went wrong")
}

func (app *application) notFound(c echo.Context) error {
	app.logger.Warnw("not found", "method", c.Request().Method, "path", c.Request().URL.Path)
	return writeJSONError(c, http.StatusNotFound, "not found")

}

func (app *application) conflictError(c echo.Context, err error) error {
	app.logger.Errorw("conflict error", "error", err.Error(), "method", c.Request().Method, "path", c.Request().URL.Path)
	return writeJSONError(c, http.StatusConflict, err.Error())
}

func (app *application) badRequest(c echo.Context, err error) error {
	app.logger.Warnw("bad request", "error", err.Error(), "method", c.Request().Method, "path", c.Request().URL.Path)
	return writeJSONError(c, http.StatusBadRequest, err.Error())
}

func (app *application) unauthorized(c echo.Context, err error) error {
	app.logger.Errorw("unauthorized", "method", c.Request().Method, "path", c.Request().URL.Path, "error", err.Error())
	return writeJSONError(c, http.StatusUnauthorized, "unauthorized")
}
