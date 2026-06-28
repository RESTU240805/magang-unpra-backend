package main

import (
	"fmt"
	"os"

	"magang-unpra-backend/config"
	"magang-unpra-backend/seeders"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run cmd/create_admin.go <email> <password> [name]")
		fmt.Println("Example: go run cmd/create_admin.go admin@telpp.com mysecurepassword \"Admin Name\"")
		os.Exit(1)
	}

	email := os.Args[1]
	password := os.Args[2]
	name := "Administrator"
	if len(os.Args) >= 4 {
		name = os.Args[3]
	}

	if len(password) < 8 {
		fmt.Println("Password must be at least 8 characters")
		os.Exit(1)
	}

	config.ConnectDB()
	seeders.CreateAdmin(email, password, name)
}
