package main

import (
	"database/sql"
	"log"

	"github.com/Rich-dev-team/stock-prediction-backend/api"
	db "github.com/Rich-dev-team/stock-prediction-backend/db/sqlc"
	"github.com/Rich-dev-team/stock-prediction-backend/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connec to db", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
