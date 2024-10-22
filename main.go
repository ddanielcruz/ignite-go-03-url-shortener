package main

import (
	"database/sql"
	"log/slog"
	"os"

	_ "modernc.org/sqlite"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	db, err := sql.Open("sqlite", "file:sqlite.db")
	if err != nil {
		slog.Error("failed to open database", "error", err)
		os.Exit(1)
	}

	// Close the database connection when the program exits
	defer db.Close()

	// Create the users table if it doesn't exist
	createTableSql := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);
	`

	_, err = db.Exec(createTableSql)
	if err != nil {
		slog.Error("failed to create table", "error", err)
		os.Exit(1)
	}

	// Insert a user into the database
	insertUserSql := `
	INSERT INTO users (name) VALUES (?)
	`

	res, err := db.Exec(insertUserSql, "Daniel Cruz")
	if err != nil {
		slog.Error("failed to insert user", "error", err)
		os.Exit(1)
	}

	// Query the created user
	queryUserSql := `
	SELECT * FROM users WHERE id = ?
	`

	userId, _ := res.LastInsertId()
	var user User

	err = db.QueryRow(queryUserSql, userId).Scan(&user.ID, &user.Name)
	if err != nil {
		slog.Error("failed to query user", "error", err)
		os.Exit(1)
	}

	slog.Info("user", "user", user)
}
