# Ignite Go 03 - URL Shortener

This project was created as a learning exercise to explore various storage types and database-related technologies in Go. The main focus was on understanding and implementing different SQL databases and Redis. The final project is a URL shortener API that uses Redis as the primary storage solution.

## Technologies Explored

- SQLite
- MySQL
- PostgreSQL
- sqlc (SQL Compiler)
- tern (SQL migration tool)
- sqlx (Extensions to database/sql)
- Redis

## Getting Started

To run this project, you'll need Docker and Go installed on your system.

1. Clone the repository:

   ```
   git clone https://github.com/ddanielcruz/ignite-go-03-url-shortener.git
   cd ignite-go-03-url-shortener
   ```

2. Start the required services using Docker Compose:

   ```
   docker-compose up -d
   ```

3. Run the project:
   ```
   go run cmd/api/main.go
   ```

The API should now be running and accessible at `http://localhost:8080` (or whichever port you've configured).
