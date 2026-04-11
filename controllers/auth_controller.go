package controllers

import (
	"marketplace/configs"
	"marketplace/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret")

func Register(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	config.DB.Create(&user)
	c.JSON(201, gin.H{"message": "User registered"})
}

func Login(c *gin.Context) {
	var input models.User
	var user models.User

	c.BindJSON(&input)

	if err := config.DB.Where("email = ? AND password = ?", input.Email, input.Password).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, _ := token.SignedString(jwtKey)

	c.JSON(200, gin.H{
		"token": tokenString,
	})
}
