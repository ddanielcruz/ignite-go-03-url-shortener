package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"sql/internal/api"
	"sql/internal/store"

	"github.com/redis/go-redis/v9"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to run", "error", err)
		os.Exit(1)
	}

	slog.Info("server stopped")
}

func run() error {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	store := store.NewStore(rdb)
	handler := api.NewHandler(store)

	server := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	slog.Info("starting server", "address", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
