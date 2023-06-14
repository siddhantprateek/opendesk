package main

import (
	"log"

	apiGateway "github.com/siddhantprateek/opendesk/api_gateway/routes"
)

func main() {

	err := apiGateway.Init()
	if err != nil {
		log.Fatal("Unable to Start Opendesk API Gateway.")
	}
}
