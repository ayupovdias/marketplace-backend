package routes

import (
	"notification-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/health", handlers.HealthCheck)

	r.POST("/notify", handlers.Notify)
}
