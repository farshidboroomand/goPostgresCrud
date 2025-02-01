package main

import (
	"log"

	"github.com/farshidboroomand/goPostgresCrud/internal/database"
	"github.com/farshidboroomand/goPostgresCrud/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to initialize Dataabse Client: %s", err)
	}
	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatalf(err.Error())
	}
}
