package main

import (
	"log"
	"net/http"
)

type config struct {
	addr string
}

type application struct {
	config config
}

func (app *application) run() error {
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
	}

	log.Printf("server started at %s", app.config.addr)
	return srv.ListenAndServe()
}
