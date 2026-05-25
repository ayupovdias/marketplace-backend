package main

import (
	"log"

	"notification-service/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	routes.SetupRoutes(r)

	log.Println("Notification service running on port 8082")

	if err := r.Run(":8082"); err != nil {
		log.Fatal(err)
	}
}
