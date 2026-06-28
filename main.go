package main

import (
	"log"
	"os"

	"magang-unpra-backend/config"
	"magang-unpra-backend/routes"
	"magang-unpra-backend/seeders"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" || jwtSecret == "change_this_to_a_random_secret_key" {
		log.Fatal("JWT_SECRET must be set to a strong random value in .env")
	}

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "release"
	}
	gin.SetMode(ginMode)

	config.ConnectDB()
	seeders.SeedAdmin()
	seeders.SeedMenus()
	seeders.SeedCompanyProfile()
	seeders.SeedCreeds()
	seeders.SeedCompanyDocuments()
	seeders.SeedCommunityCards()
	seeders.SeedPulpProcessSections()
	seeders.SeedPulpProcessRecoveries()
	seeders.SeedSafetyPolicies()
	seeders.SeedSafetyK3Targets()
	seeders.SeedSafetyK3Programs()
	seeders.SeedSafetySliders()
	seeders.SeedSupplyChainStrategies()
	seeders.SeedSupplyChainSustainabilityItems()
	seeders.SeedSupplyChainPolicies()
	seeders.SeedCsrVisionStrategies()
	seeders.SeedCsrReports()

	r := routes.SetupRoutes()
	r.MaxMultipartMemory = 10 << 20

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
