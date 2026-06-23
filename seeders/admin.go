package seeders

import (
	"log"
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"

	"golang.org/x/crypto/bcrypt"
)

func SeedAdmin() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password:", err)
	}

	var existingUser models.User
	result := config.DB.Where("email = ?", "admin@telpp.com").First(&existingUser)
	if result.Error == nil {
		config.DB.Model(&existingUser).Update("password", string(hashedPassword))
		log.Println("Admin password updated: admin@telpp.com / admin123")
		return
	}

	admin := models.User{
		Name:     "Administrator",
		Email:    "admin@telpp.com",
		Password: string(hashedPassword),
		Role:     "admin",
	}

	config.DB.Create(&admin)
	log.Println("Admin seeded: admin@telpp.com / admin123")
}
