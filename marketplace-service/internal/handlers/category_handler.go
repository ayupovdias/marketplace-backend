package handlers

import (
	"marketplace-service/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	DB *gorm.DB
}

func (h *CategoryHandler) Create(c *gin.Context) {

	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.DB.Create(&category).Error

	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to create category",
		})
		return
	}

	c.JSON(201, gin.H{
		"category": category,
	})
}

func (h *CategoryHandler) GetAll(c *gin.Context) {

	var categories []models.Category

	h.DB.Find(&categories)

	c.JSON(200, gin.H{
		"categories": categories,
	})
}
func (h *CategoryHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	var category models.Category

	if err := h.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "category not found",
		})
		return
	}

	c.JSON(http.StatusOK, category)
}
func (h *CategoryHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var category models.Category

	if err := h.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "category not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.DB.Save(&category)

	c.JSON(http.StatusOK, category)
}
func (h *CategoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.DB.Delete(&models.Category{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete category",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "category deleted",
	})
}
