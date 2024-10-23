package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type store struct {
	rdb *redis.Client
}

type Store interface {
	SaveShortenedURL(ctx context.Context, _url string) (string, error)
	GetShortenedURL(ctx context.Context, code string) (string, error)
}

func NewStore(rdb *redis.Client) Store {
	return store{rdb: rdb}
}

// Use _url instead of url to avoid confusion with the url package.
func (s store) SaveShortenedURL(ctx context.Context, _url string) (string, error) {
	var code string
	var success bool
	const attempts = 5

	// Try to generate a code that is not already in the database.
	for range attempts {
		code = genCode()
		if err := s.rdb.HGet(ctx, "shortened_urls", code).Err(); err != nil {
			// If record does not exist, we can use this code.
			if errors.Is(err, redis.Nil) {
				success = true
				break
			}

			return "", fmt.Errorf("failed to get code: %w", err)
		}
	}

	// Check if we were able to generate a code that is not in the database.
	if !success {
		return "", errors.New("failed to generate code")
	}

	// Save the code and url in the database.
	if err := s.rdb.HSet(ctx, "shortened_urls", code, _url).Err(); err != nil {
		return "", fmt.Errorf("failed to set code: %w", err)
	}

	return code, nil
}

func (s store) GetShortenedURL(ctx context.Context, code string) (string, error) {
	_url, err := s.rdb.HGet(ctx, "shortened_urls", code).Result()

	if err != nil {
		return "", fmt.Errorf("failed to get url: %w", err)
	}

	return _url, nil
}
