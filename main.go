package main

import (
	"database/sql"
	"github.com/Srimanth-Reddy/simplebank/api"
	"github.com/Srimanth-Reddy/simplebank/db/sqlc"
	"github.com/Srimanth-Reddy/simplebank/util"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}

	store := sqlc.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
}
