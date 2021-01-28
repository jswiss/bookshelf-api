package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jswiss/bookshelf/config"
)

// "DBConnection - connect to database"
func DbConnection() {
	fmt.Printf("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	dbName, dbUser, dbPass := config.LocalEnv()

	dbURL := fmt.Sprintf("postgresql://%s:%s@db/%s", dbUser, dbPass, dbName)
	fmt.Println("DBURL: %s", dbURL)

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connection to database: %s failed because: %v\n", dbName, err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

}
