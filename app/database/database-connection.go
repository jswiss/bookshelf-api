package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jswiss/bookshelf/config"
)

// DB ...data
var DB *sql.DB

// Connect ...
func Connect() (*sql.DB, error) {
	var err error
	dbName, dbUser, dbPass := config.LocalEnv()

	dbURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "db", 5432, dbUser, dbPass, dbName)

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connection to database: %s failed because: %v\n", dbName, err)
		os.Exit(1)
	}

	if err = conn.Ping(); err != nil {
		return conn, err
	}

	return conn, err
}
