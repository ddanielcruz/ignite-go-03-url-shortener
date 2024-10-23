package api

import (
	"net/http"
	"sql/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func NewHandler(store store.Store) http.Handler {
	r := chi.NewMux()

	// Middleware
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	// Routes
	r.Route("/url", func(r chi.Router) {
		r.Post("/shorten", handleShortenURL(store))
		r.Get("/{code}", handleGetShortenedURL(store))
	})

	return r
}
