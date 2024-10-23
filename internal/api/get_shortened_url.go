package api

import (
	"errors"
	"log/slog"
	"net/http"
	"sql/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

type getShortenedURLResponse struct {
	FullURL string `json:"full_url"`
}

func handleGetShortenedURL(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		fullURL, err := store.GetShortenedURL(r.Context(), code)

		if err != nil {
			if errors.Is(err, redis.Nil) {
				sendJSON(w, Response{Error: "not found"}, http.StatusNotFound)
				return
			}

			sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
			slog.Error("failed to get shortened URL", "error", err)
			return
		}

		sendJSON(w, Response{Data: getShortenedURLResponse{FullURL: fullURL}}, http.StatusOK)
	}
}
