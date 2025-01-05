package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {
	e := echo.New()
	r := e.Group("/v1")

	r.GET("/ping", app.pingHandler)

	return e
}

func (app *application) run(mux http.Handler) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Minute,
	}

	return srv.ListenAndServe()
}
