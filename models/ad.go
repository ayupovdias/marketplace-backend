package models

import "gorm.io/gorm"

type Ad struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	City        string `json:"city"`
}
