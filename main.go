package main

import (
	"fmt"
	"go-gamemaster-api/database"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvs()
	database.Connect()
	appRouter := gin.Default()
	appRouter.Run(":4444")
	fmt.Println("Server is running on port 4444")
}

func loadEnvs() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}
