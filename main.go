package main

import (
	"log"
	"os"

	"magang-unpra-backend/config"
	"magang-unpra-backend/routes"
	"magang-unpra-backend/seeders"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()
	seeders.SeedAdmin()
	seeders.SeedMenus()
	seeders.SeedCompanyProfile()
	seeders.SeedCreeds()
	seeders.SeedCompanyDocuments()
	seeders.SeedCommunityCards()

	r := routes.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
