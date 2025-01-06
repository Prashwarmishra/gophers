package main

import "gophers/internal/env"

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8000"),
	}

	app := &application{
		config: cfg,
	}

	app.run(app.mount())
}