package main

import (
	"log"
	"marketplace-service/internal/database"
	"marketplace-service/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("failed to load env")
	}

	db := database.ConnectDB()

	r := gin.Default()

	routes.SetupRoutes(r, db)

	log.Println("Marketplace service running on :8081")

	r.Run(":8081")
}
