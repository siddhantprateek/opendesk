package main

import (
	"log"

	connection "github.com/siddhantprateek/opendesk/storage/connection"
	app "github.com/siddhantprateek/opendesk/storage/routes"
)

func main() {
	// MongoDB Database connection.
	connection.MongoDBdatabase()

	// Application API entrypoint
	err := app.Init()
	if err != nil {
		log.Fatal("Unable to start Opendesk Storage Microserivce.")
	}
}
