package main

import (
	"calibri/cmd/api"
	"calibri/internal/database"
	"log"
)

func main() {
	db, err := database.PostgresConnect()
	if err != nil {
		log.Fatalln(err)
	}

	server := api.NewApiServer(":8070", db)
	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}

}
