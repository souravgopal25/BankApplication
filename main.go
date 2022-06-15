package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/souravgopal25/BankApplication/api"
	db "github.com/souravgopal25/BankApplication/db/sqlc"
	"github.com/souravgopal25/BankApplication/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalln("Cannot Load config file", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("Cannot create to db: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot Start Server", err)
	}

}
