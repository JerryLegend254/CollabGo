package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/JerryLegend254/CollabGo/internal/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type userContextKey string

var userCtxKey userContextKey = "user"

func (app *application) LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		// Call the next handler
		err := next(c)
		if err != nil {
			// Log error details if any
			app.logger.Errorw(
				"Request Error",
				"error", err.Error(),
				"method", c.Request().Method,
				"uri", c.Request().URL.Path,
			)
		}

		// Log response status and duration
		duration := time.Since(start)
		app.logger.Infow(
			"Completed Request",
			"method", c.Request().Method,
			"uri", c.Request().URL.Path,
			"status", c.Response().Status,
			"duration", duration.String(),
		)

		return err
	}
}

func extractAccessTokenFromToken(token *jwt.Token) (string, error) {
	// Ensure the claims are of type jwt.MapClaims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if accessToken, exists := claims["access_token"].(string); exists {
			return accessToken, nil
		}
		return "", fmt.Errorf("access_token not found in claims")
	}

	return "", fmt.Errorf("invalid token claims")
}

func (app *application) AuthTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return app.unauthorized(c, fmt.Errorf("authorization header is missing"))
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return app.unauthorized(c, fmt.Errorf("authorization header is missing or invalid"))
		}

		token := parts[1]
		_, err := app.auth.JWT.ValidateToken(token)
		if err != nil {
			return app.unauthorized(c, err)
		}

		oauthToken, err := app.auth.JWT.ParseJWTToken(token)
		if err != nil {
			return app.unauthorized(c, err)
		}

		ctx := c.Request().Context()

		user, err := app.getUser(ctx, oauthToken)
		if err != nil {
			return app.unauthorized(c, err)
		}

		ctx = context.WithValue(ctx, "user", user)

		fmt.Println("user", user)

		req := c.Request().WithContext(ctx)
		c.SetRequest(req)
		return next(c)
	}
}
func (app *application) getUser(ctx context.Context, token *oauth2.Token) (*auth.User, error) {
	client := app.config.auth.oauth.config.Client(context.Background(), token)
	return auth.GetSpotifyUserInfo(client)
}
