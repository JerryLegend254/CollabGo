package store

import (
	"context"
	"database/sql"
	"time"
)

var (
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Playlists interface {
		Create(context.Context, *Playlist) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Playlists: &PlaylistStore{db},
	}
}
