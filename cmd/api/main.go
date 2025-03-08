package main

import (
	"github.com/JerryLegend254/CollabGo/internal/auth"
	"github.com/JerryLegend254/CollabGo/internal/db"
	"github.com/JerryLegend254/CollabGo/internal/env"
	"github.com/JerryLegend254/CollabGo/internal/logger"
	"github.com/JerryLegend254/CollabGo/internal/store"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:           env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/collabgo?sslmode=disable"),
			maxOpenConns:   env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns:   env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTimeout: env.GetString("DB_MAX_IDLE_TIMEOUT", "15m"),
		},
		auth: authConfig{
			oauth: oauthConfig{
				config: &oauth2.Config{
					ClientID:     env.GetString("SPOTIFY_CLIENT_ID", ""),
					ClientSecret: env.GetString("SPOTIFY_CLIENT_SECRET", ""),
					Endpoint:     spotify.Endpoint,
					RedirectURL:  "http://localhost:8080/v1/authentication/callback/spotify",
					Scopes:       []string{"user-read-email", "user-read-recently-played"},
				},
				state: auth.GenerateRandomState(),
			},
			token: tokenConfig{
				secret: env.GetString("TOKEN_SECRET", "secret"),
				iss:    "collabgo",
			},
		},
	}

	// Logger setup
	logger := logger.NewLogger()

	defer logger.Sync()

	// Database
	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTimeout)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Database connection successful")

	defer db.Close()
	store := store.NewStorage(db)

	tokenHost := "collabgo"
	authenticator := auth.NewAuthenticator(cfg.auth.oauth.config, cfg.auth.oauth.state, cfg.auth.token.secret, tokenHost, tokenHost)

	app := &application{
		config: cfg,
		store:  store,
		auth:   authenticator,
		logger: logger,
	}

	mux := app.mount()

	logger.Infof("server listening at %s", app.config.addr)
	logger.Fatal(app.run(mux))

}
