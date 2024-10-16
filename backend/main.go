package main

import (
	"backend/config"
	"backend/routes"
	"backend/database"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()
	database.Connect()
	database.Seed()

	r := routes.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}