package controllers

import (
	"marketplace/configs"
	"marketplace/models"
	"marketplace/payloads"
	"marketplace/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input payloads.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"errors": utils.FormatValidationError(err),
		})
		return
	}

	if input.Password != input.ConfirmPassword {
		c.JSON(400, gin.H{
			"error": "Passwords do not match",
		})
		return
	}

	user := models.User{
		Username:    input.Username,
		PhoneNumber: input.PhoneNumber,
		Email:       input.Email,
		Password:    input.Password,
	}

	config.DB.Create(&user)

	c.JSON(201, gin.H{
		"message": "User registered",
	})
}
func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Where("email = ? AND password = ?", input.Email, input.Password).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
