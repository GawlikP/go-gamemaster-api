package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() {
	var err error
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	name := os.Getenv("DATABASE_NAME")
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	conn := fmt.Sprintf(`host=%s port=%s dbname=%s user=%s password=%s sslmode=disable
		TimeZone=Europe/Warsaw`, host, port, name, username, password)

	DB, err = gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database!")
	}
}
