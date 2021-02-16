package main

import (
	"database/sql"
	"log"

	db "github.com/jswiss/bookshelf/app/database/sqlc"
	api "github.com/jswiss/bookshelf/app/server"
	"github.com/jswiss/bookshelf/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, *store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(":3000")
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
