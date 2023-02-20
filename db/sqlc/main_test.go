package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/domingo1021/golang-bank-account/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../ci_test")
	if err != nil {
		log.Fatal("cannot load env variable: ", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
