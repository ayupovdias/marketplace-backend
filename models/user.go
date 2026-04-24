package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `json:"username" binding:"required,min=3"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email" binding:"required,email" gorm:"unique"`
	Password    string `json:"password" binding:"required,min=4"`
}
