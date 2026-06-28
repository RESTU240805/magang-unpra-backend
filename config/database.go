package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"magang-unpra-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "disable"
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get sql.DB:", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if os.Getenv("GIN_MODE") != "release" || os.Getenv("AUTO_MIGRATE") == "true" {
		db.AutoMigrate(
			&models.User{},
			&models.News{},
			&models.NewsImage{},
			&models.Product{},
			&models.ProductImage{},
			&models.ProductSlider{},
			&models.Creed{},
			&models.Menu{},
			&models.CompanyProfile{},
			&models.CompanyDocument{},
			&models.ProductPage{},
			&models.AboutSection{},
			&models.CommunityCard{},
			&models.CommunityCardImage{},
			&models.TeamMember{},
			&models.OrgChart{},
			&models.OrgGroup{},
			&models.OrgNode{},
			&models.ForestWoodType{},
			&models.ForestApproach{},
			&models.ForestSlider{},
			&models.PeopleDevelopmentPage{},
			&models.PeopleDevelopmentPillar{},
			&models.PeopleDevelopmentSlider{},
			&models.PulpProcessSection{},
			&models.PulpProcessRecovery{},
			&models.SafetyPolicy{},
			&models.SafetyK3Target{},
			&models.SafetyK3Program{},
			&models.SafetySlider{},
			&models.SupplyChainStrategy{},
			&models.SupplyChainSustainabilityItem{},
			&models.SupplyChainPolicy{},
			&models.CsrVisionContent{},
			&models.CsrVisionStrategy{},
			&models.CsrReport{},
			&models.ContactInfo{},
			&models.ContactOffice{},
		)
		log.Println("Database migrated successfully!")
	}

	log.Println("Database connected successfully!")
	DB = db
}
