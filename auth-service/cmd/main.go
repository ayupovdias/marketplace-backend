package main

import (
	"log"

	"auth-service/internal/config"
	"auth-service/internal/database"
	"auth-service/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	db := database.Connect(cfg)

	router := gin.Default()

	routes.SetupRoutes(router, db)

	log.Println("Server running on port", cfg.Port)

	router.Run(":" + cfg.Port)
}
