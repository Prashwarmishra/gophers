package main

import (
	"gophers/internal/env"
	"gophers/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8000"),
	}

	app := &application{
		config: cfg,
		store: store.NewStorage(nil),
	}

	app.run(app.mount())
}