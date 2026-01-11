package main

import (
	"log"

	"github.com/rissabekov-wes/social/internal/env"
	"github.com/rissabekov-wes/social/internal/store"
)

const version = "1.0.0"

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8081"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://user:pass@localhost:5432/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 25),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 25),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	log.Printf("Config: %+v", cfg)

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))

}
