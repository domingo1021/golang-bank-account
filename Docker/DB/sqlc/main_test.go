package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// const variabe of db setup will be moved to env variable latter.
const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password1234@localhost:5432/bank_account?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: " , err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}