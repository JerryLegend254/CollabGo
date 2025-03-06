package main

import (
	"github.com/labstack/echo/v4"
)

func (app *application) LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// log the request
		app.logger.Infow(
			"INFO",
			"method", c.Request().Method,
			"uri", c.Request().URL.Path,
			"query", c.Request().URL.RawQuery,
		)

		// call the next middleware/handler
		err := next(c)
		if err != nil {
			app.logger.Errorw(
				"ERROR",
				"error", err.Error(),
			)
			return err
		}

		return nil
	}
}
