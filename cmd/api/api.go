package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (a *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health-check", a.healthCheckHandler)
	})

	return r
}

func (app *application) run(mux http.Handler) {
	server := http.Server {
		Addr: app.config.addr,
		Handler: mux,
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout: time.Minute,
	}
	
	log.Printf("Server running on port: %v\n", app.config.addr)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("err in initializing server:", err)
		panic("failed to initialize server")
	}
}
