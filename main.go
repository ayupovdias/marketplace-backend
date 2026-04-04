package main

import (
	"marketplace/configs"
	"marketplace/models"
	"marketplace/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()
	config.DB.AutoMigrate(&models.Ad{})

	routes.SetupRoutes(r)

	r.Run(":8080")
}
