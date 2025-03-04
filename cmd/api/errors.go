package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *application) internalServerError(c echo.Context, err error) error {
	fmt.Println("internal server error", "error", err.Error(), "method", c.Request().Method, "path", c.Request().URL.Path)
	return writeJSONError(c, http.StatusInternalServerError, "something went wrong")
}

func (app *application) notFound(c echo.Context) error {
	fmt.Println("not found", "method", c.Request().Method, "path", c.Request().URL.Path)
	return writeJSONError(c, http.StatusNotFound, "not found")

}

func (app *application) conflictError(c echo.Context, err error) error {
	fmt.Println("conflict error", "error", err.Error(), "method", c.Request().Method, "path", c.Request().URL.Path)
	return writeJSONError(c, http.StatusConflict, err.Error())
}

func (app *application) badRequest(c echo.Context, err error) error {
	fmt.Println("bad request", "error", err.Error(), "method", c.Request().Method, "path", c.Request().URL.Path)
	return writeJSONError(c, http.StatusBadRequest, err.Error())
}
