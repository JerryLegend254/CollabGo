package main

import (
	"net/http"

	"github.com/JerryLegend254/CollabGo/internal/store"
	"github.com/labstack/echo/v4"
)

type CreatePlaylistPayload struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"max=255"`
	IsPublic    bool   `json:"is_public" binding:"required,boolean"`
}

func (app *application) createPlaylistHandler(c echo.Context) error {
	var payload CreatePlaylistPayload

	if err := readJSON(c, &payload); err != nil {
		return app.badRequest(c, err)
	}

	if err := Validate.Struct(&payload); err != nil {
		return app.badRequest(c, err)
	}

	playlist := &store.Playlist{
		Name:        payload.Name,
		IsPublic:    payload.IsPublic,
		Description: payload.Description,
	}
	err := app.store.Playlists.Create(c.Request().Context(), playlist)
	if err != nil {
		return app.internalServerError(c, err)
	}

	return app.jsonResponse(c, http.StatusOK, playlist)
}
