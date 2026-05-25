package handlers

import (
	"marketplace-service/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FavoriteHandler struct {
	DB *gorm.DB
}

func (h *FavoriteHandler) Add(c *gin.Context) {

	adID, _ := strconv.Atoi(c.Param("ad_id"))

	userID := c.MustGet("user_id").(uint)

	favorite := models.Favorite{
		UserID:          userID,
		AdvertisementID: uint(adID),
	}

	err := h.DB.Create(&favorite).Error

	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to add favorite",
		})
		return
	}

	c.JSON(201, gin.H{
		"favorite": favorite,
	})
}

func (h *FavoriteHandler) GetAll(c *gin.Context) {

	userID := c.MustGet("user_id").(uint)

	var favorites []models.Favorite

	h.DB.
		Where("user_id = ?", userID).
		Find(&favorites)

	c.JSON(200, gin.H{
		"favorites": favorites,
	})
}
func (h *FavoriteHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	var favorite models.Favorite

	if err := h.DB.First(&favorite, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "favorite not found",
		})
		return
	}

	c.JSON(http.StatusOK, favorite)
}
func (h *FavoriteHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.DB.Delete(&models.Favorite{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "favorite deleted",
	})
}
