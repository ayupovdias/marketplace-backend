package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
