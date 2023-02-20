package main

import (
	"database/sql"
	"log"

	"github.com/domingo1021/golang-bank-account/api"
	db "github.com/domingo1021/golang-bank-account/db/sqlc"
	"github.com/domingo1021/golang-bank-account/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".", "app")
	if err != nil {
		log.Fatal("cannot read env variable: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("start web server failed: ", err)
	}
}
