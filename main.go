package database

import (
	"fmt"
	"log"

	"github.com/jswiss/bookshelf/app/database"
	api "github.com/jswiss/bookshelf/app/server"
	_ "github.com/lib/pq"
	db "github.com/techschool/simplebank/db/sqlc"
)

func initDatabase() {
	var err error
	database.Connect()
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}

func main() {
	initDatabase()
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
