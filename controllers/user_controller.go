package controllers

import (
	"marketplace/configs"
	"marketplace/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	var user models.User
	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, user)
}
func UpdateUser(c *gin.Context) {
	var user models.User

	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.ShouldBindJSON(&user)
	config.DB.Save(&user)

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	config.DB.Delete(&models.User{}, c.Param("id"))
	c.JSON(200, gin.H{"message": "User deleted"})
}
