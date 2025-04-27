package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver import
)

// DB holds the database connection pool.
var DB *sql.DB

// InitializeDB initializes the database connection.
// Update the DSN (Data Source Name) with your actual database credentials.
// Recommended format: "username:password@tcp(host:port)/dbname"
func InitializeDB() {
	var err error

	// TODO: Replace with configuration or environment variables for production use.
	DSN := "user:password@tcp(localhost:3306)/dbname"

	DB, err = sql.Open("mysql", DSN)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Ping ensures a live database connection.
	if err = DB.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Println("Database successfully connected!")
}
