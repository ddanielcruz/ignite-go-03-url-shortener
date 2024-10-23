package store

import "golang.org/x/exp/rand"

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func genCode() string {
	const length = 8
	bytes := make([]byte, length)

	for i := range bytes {
		bytes[i] = characters[rand.Intn(len(characters))]
	}

	return string(bytes)
}
