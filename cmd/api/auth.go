package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/JerryLegend254/CollabGo/internal/auth"
	"github.com/labstack/echo/v4"
)

func (app *application) handleSpotifyLoginHandler(c echo.Context) error {
	url := app.config.auth.oauth.config.AuthCodeURL(app.config.auth.oauth.state)
	http.Redirect(c.Response().Writer, c.Request(), url, http.StatusTemporaryRedirect)
	return nil
}

func (app *application) handleCallbackHandler(c echo.Context) error {
	state := c.QueryParam("state")
	code := c.QueryParam("code")
	if state != app.config.auth.oauth.state {
		http.Error(c.Response().Writer, "Invalid state", http.StatusBadRequest)
		return errors.New("Invalid state")
	}

	token, err := app.config.auth.oauth.config.Exchange(context.Background(), code)
	if err != nil {
		http.Error(c.Response().Writer, "Failed to exchange token", http.StatusInternalServerError)
		app.logger.Error("Token exchange error:", err)
		return errors.New("Failed to exchange token")
	}

	client := app.config.auth.oauth.config.Client(context.Background(), token)
	userInfo, err := auth.GetSpotifyUserInfo(client)
	if err != nil {
		http.Error(c.Response().Writer, "Failed to get user info", http.StatusInternalServerError)
		app.logger.Error("User info error:", err)
		return errors.New("Failed to get user info")
	}
	fmt.Fprintf(c.Response().Writer, "Logged in successfully! User: %s\n", userInfo)

	return nil
}
