package main

import (
	"database/sql"
	"log"

	"github.com/domingo1021/golang-bank-account/api"
	db "github.com/domingo1021/golang-bank-account/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:password1234@localhost:5432/bank_account?sslmode=disable"
	serverAddress = "localhost:8080"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("start web server failed: ", err)
	}
}
