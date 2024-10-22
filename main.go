package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	db, err := sql.Open("mysql", "root:root@/rocketseat")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	query := `CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL
	)`

	if _, err := db.Exec(query); err != nil {
		panic(err)
	}

	fmt.Println("Table created successfully")

	query = `INSERT INTO users (name) VALUES (?)`
	input := fmt.Sprintf("User %d", time.Now().UnixMilli())

	if _, err := db.Exec(query, input); err != nil {
		panic(err)
	}

	fmt.Println("User created successfully")

	query = `SELECT * FROM users`
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name)
		fmt.Printf("User: %+v\n", user)

	}

	fmt.Println("Users fetched successfully")
}
