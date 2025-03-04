package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func readJSON(c echo.Context, data interface{}) error {
	maxBytes := int64(1 << 20) // 1 MB
	c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, maxBytes)

	dec := json.NewDecoder(c.Request().Body)
	dec.DisallowUnknownFields()

	return dec.Decode(data)
}

func writeJSONError(c echo.Context, status int, message string) error {
	type errorResponse struct {
		Error string `json:"error"`
	}

	data := &errorResponse{
		Error: message,
	}

	return c.JSON(status, data)
}

func (app *application) jsonResponse(c echo.Context, status int, data interface{}) error {
	type jsonResponse struct {
		Data interface{} `json:"data"`
	}
	return c.JSON(status, &jsonResponse{Data: data})
}
