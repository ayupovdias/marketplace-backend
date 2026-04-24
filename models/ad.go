package models

import "gorm.io/gorm"

type Ad struct {
	gorm.Model
	Title       string `json:"title" binding:"required,min=3"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required,gt=0"`
	City        string `json:"city" binding:"required"`

	UserID uint `json:"user_id"`
}
