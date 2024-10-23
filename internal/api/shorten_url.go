package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"sql/internal/store"
)

type shortenURLRequest struct {
	URL string `json:"url"`
}

type shortenURLResponse struct {
	Code string `json:"code"`
}

func handleShortenURL(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body shortenURLRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, Response{Error: "invalid request body"}, http.StatusUnprocessableEntity)
			return
		}

		parsedURL, ok := isURL(body.URL)
		if !ok {
			sendJSON(w, Response{Error: "invalid URL"}, http.StatusBadRequest)
			return
		}

		code, err := store.SaveShortenedURL(r.Context(), parsedURL.String())
		if err != nil {
			sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
			slog.Error("failed to save shortened URL", "error", err)
			return
		}

		sendJSON(w, Response{Data: shortenURLResponse{Code: code}}, http.StatusCreated)
	}
}
