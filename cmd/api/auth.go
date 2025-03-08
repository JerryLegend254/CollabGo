package main

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

	claims := jwt.MapClaims{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"exp":           token.Expiry.Unix(),
		"iat":           time.Now().Unix(),
		"iss":           app.config.auth.token.iss,
		"aud":           app.config.auth.token.iss,
	}

	jwtToken, err := app.auth.JWT.GenerateToken(claims)
	if err != nil {
		return app.internalServerError(c, err)
	}

	return app.jsonResponse(c, http.StatusOK, jwtToken)
}
