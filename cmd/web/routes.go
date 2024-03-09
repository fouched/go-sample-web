package main

import (
	"github.com/fouched/go-sample-web/internal/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)

	return mux
}
