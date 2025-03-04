package store

import (
	"context"
	"database/sql"
)

type PlaylistStore struct {
	db *sql.DB
}

type Playlist struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsPublic    bool   `json:"is_public"`
}

// TODO: add user_id once auth is implemented

func (s *PlaylistStore) Create(ctx context.Context, playlist *Playlist) error {
	query := `INSERT INTO playlist (name, description, is_public) VALUES ($1, $2, $3) RETURNING id;`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	if err := s.db.QueryRowContext(ctx, query, &playlist.Name, &playlist.Description, &playlist.IsPublic).Scan(&playlist.ID); err != nil {
		return err
	}

	return nil
}
