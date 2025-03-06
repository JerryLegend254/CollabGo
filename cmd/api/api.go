package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JerryLegend254/CollabGo/internal/auth"
	"github.com/JerryLegend254/CollabGo/internal/store"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type application struct {
	config config
	store  store.Storage
	auth   auth.Authenticator
}

type config struct {
	addr string
	db   dbConfig
	auth authConfig
}

type dbConfig struct {
	addr           string
	maxOpenConns   int
	maxIdleConns   int
	maxIdleTimeout string
}

type authConfig struct {
	oauth oauthConfig
}

type oauthConfig struct {
	config *oauth2.Config
	state  string
}

var (
	upgrader = websocket.Upgrader{}
)

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}

func (app *application) mount() http.Handler {
	e := echo.New()

	r := e.Group("/v1")

	r.GET("/ping", app.pingHandler)
	r.GET("/ws", hello)

	authentication := r.Group("/authentication")

	authentication.GET("/login/spotify", app.handleSpotifyLoginHandler)
	authentication.GET("/callback/spotify", app.handleCallbackHandler)

	playlistsRoutes := r.Group("/playlists")
	playlistsRoutes.POST("", app.createPlaylistHandler)

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
