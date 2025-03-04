package main

import (
	"log"

	"github.com/JerryLegend254/CollabGo/internal/db"
	"github.com/JerryLegend254/CollabGo/internal/env"
	"github.com/JerryLegend254/CollabGo/internal/store"
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
	}

	// Database
	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTimeout)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	store := store.NewStorage(db)
	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Printf("server listening at %s", app.config.addr)
	log.Fatal(app.run(mux))

}
