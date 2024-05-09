package main

import (
	"log"
	"main/config"
	"main/router"
)

func main() {
	// Init DB connection
	err := config.DatabaseConnection()
	if err != nil {
		log.Fatal("Error conecting to the DB")
	}

	// SetUp the router
	r := router.SetupRouter()

	r.Run()
}
