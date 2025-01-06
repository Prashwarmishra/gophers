package main

import (
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (a *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthCheckHandler)

	return mux
}

func (a *application) run(mux *http.ServeMux) {
	server := http.Server{
		Addr: a.config.addr,
		Handler: mux,
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout: time.Minute,
	}	

	log.Printf("Server running on port: %v", a.config.addr)

	err := server.ListenAndServe()

	if err != nil {
		panic("Failed to initialize server")
	}	
}