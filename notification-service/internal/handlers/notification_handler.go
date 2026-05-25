package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationInput struct {
	Message string `json:"message"`
	Title   string `json:"title"`
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "notification service running",
	})
}

func Notify(c *gin.Context) {

	var input NotificationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Println("========== NEW NOTIFICATION ==========")
	log.Println("TITLE:", input.Title)
	log.Println("MESSAGE:", input.Message)
	log.Println("======================================")

	c.JSON(http.StatusOK, gin.H{
		"status": "notification received",
	})
}
