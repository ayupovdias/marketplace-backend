package controllers

import (
	"marketplace/configs"
	"marketplace/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAds(c *gin.Context) {
	var ads []models.Ad
	config.DB.Find(&ads)
	c.JSON(http.StatusOK, ads)
}

func GetAd(c *gin.Context) {
	var ad models.Ad
	if err := config.DB.First(&ad, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Ad not found"})
		return
	}
	c.JSON(200, ad)
}

func CreateAd(c *gin.Context) {
	var ad models.Ad
	c.BindJSON(&ad)
	if ad.Price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "price is negative"})
		return
	}
	config.DB.Create(&ad)
	c.JSON(201, ad)
}

func UpdateAd(c *gin.Context) {
	var ad models.Ad
	if err := config.DB.First(&ad, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}
	if ad.Price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "price is negative"})
		return
	}
	c.BindJSON(&ad)
	config.DB.Save(&ad)

	c.JSON(200, ad)
}

func DeleteAd(c *gin.Context) {
	config.DB.Delete(&models.Ad{}, c.Param("id"))
	c.JSON(200, gin.H{"message": "Deleted"})
}
func GetByCity(c *gin.Context) {
	var ads []models.Ad
	config.DB.Where("city = ?", c.Param("city")).Find(&ads)
	c.JSON(200, ads)
}

func GetByPrice(c *gin.Context) {
	var ads []models.Ad
	min := c.Param("min")
	max := c.Param("max")

	config.DB.Where("price BETWEEN ? AND ?", min, max).Find(&ads)
	c.JSON(200, ads)
}

func SearchAds(c *gin.Context) {
	var ads []models.Ad
	query := c.Query("q")

	config.DB.Where("title ILIKE ?", "%"+query+"%").Find(&ads)
	c.JSON(200, ads)
}

func UpdatePrice(c *gin.Context) {
	var ad models.Ad
	config.DB.First(&ad, c.Param("id"))

	var body struct {
		Price int
	}
	c.BindJSON(&body)
	if ad.Price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "price is negative"})
		return
	}
	ad.Price = body.Price
	config.DB.Save(&ad)

	c.JSON(200, ad)
}

func LatestAds(c *gin.Context) {
	var ads []models.Ad
	config.DB.Order("created_at desc").Limit(10).Find(&ads)
	c.JSON(200, ads)
}
