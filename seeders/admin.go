package seeders

import (
	"fmt"
	"log"

	"magang-unpra-backend/config"
	"magang-unpra-backend/models"

	"golang.org/x/crypto/bcrypt"
)

func SeedAdmin() {
	var count int64
	config.DB.Model(&models.User{}).Count(&count)
	if count > 0 {
		return
	}

	log.Println("[WARNING] No admin user found in database!")
	log.Println("Create one manually:")
	log.Println("  go run cmd/create_admin.go <email> <password>")
}

func CreateAdmin(email, password, name string) {
	if email == "" || password == "" {
		log.Fatal("Email and password are required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password:", err)
	}

	var count int64
	config.DB.Model(&models.User{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		fmt.Printf("User with email '%s' already exists\n", email)
		return
	}

	if name == "" {
		name = "Administrator"
	}

	admin := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     "admin",
	}

	config.DB.Create(&admin)
	fmt.Printf("Admin created: %s\n", email)
}
