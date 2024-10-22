package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	connUrl := "postgres://postgres:postgres@localhost:5432/rocketseat"
	db, err := pgxpool.New(ctx, connUrl)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	queries := New(db)

	// Create author
	author, err := queries.CreateAuthor(ctx, CreateAuthorParams{
		Name: "Daniel",
		Bio:  pgtype.Text{String: "I'm a software engineer", Valid: true},
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Author: %+v\n", author)

	// List authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Authors: %+v\n", authors)

	// Update author
	err = queries.UpdateAuthor(ctx, UpdateAuthorParams{
		ID:   author.ID,
		Name: "Daniel",
		Bio:  pgtype.Text{String: "Am I a software engineer?", Valid: true},
	})

	if err != nil {
		panic(err)
	}

	// Get author
	author, err = queries.GetAuthor(ctx, author.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Author: %+v\n", author)

	// Delete author
	err = queries.DeleteAuthor(ctx, author.ID)
	if err != nil {
		panic(err)
	}
}
