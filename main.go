package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	connUrl := "postgres://postgres:postgres@localhost:5432/rocketseat"
	db, err := pgxpool.New(context.Background(), connUrl)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	if err := db.Ping(context.Background()); err != nil {
		panic(err)
	}

	query := `CREATE TABLE IF NOT EXISTS users (
		id BIGSERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL
	)`

	if _, err := db.Exec(context.Background(), query); err != nil {
		panic(err)
	}

	fmt.Println("Table created successfully")

	query = `INSERT INTO users (name) VALUES ($1)`
	input := fmt.Sprintf("User %d", time.Now().UnixMilli())

	if _, err := db.Exec(context.Background(), query, input); err != nil {
		panic(err)
	}

	fmt.Println("User created successfully")

	query = `SELECT * FROM users`
	rows, err := db.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			panic(err)
		}

		fmt.Printf("User: %+v\n", user)
	}

	fmt.Println("Users fetched successfully")
}
