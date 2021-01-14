package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jswiss/bookshelf/config"
	"github.com/jackc/pgx/v4"
)

func dbConnection() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	dbName, dbUser, dbPass, dbPort := config.localEnv()

	fmt.Printf("db name is %s\n", dbName)
	fmt.Printf("db user is %s\n", dbUser)
	fmt.Printf("db pass is %s\n", dbPass)
	fmt.Printf("db port is %s\n", dbPort)

	defer conn.Close(context.Background())

	var name string
	var weight int64
	err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, weight)
}
