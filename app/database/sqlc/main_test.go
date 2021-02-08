package database

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/jswiss/bookshelf/util"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig(".")
	conn, err := sql.Open(dbDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
