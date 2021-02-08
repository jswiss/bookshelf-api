package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jswiss/bookshelf/config"
	_ "github.com/lib/pq"
)

var dbUser, dbPass, dbName = config.LocalEnv()

var dbSource = fmt.Sprintf("postgresql://{}:{}@db:5432/{}?sslmode=disable", dbUser, dbPass, dbName)

const (
	dbDriver = "postgres"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
