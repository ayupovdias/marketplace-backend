package handlers

import (
	"log"
	"marketplace-service/internal/models"
	"marketplace-service/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type AdvertisementHandler struct {
	Repo *repository.AdvertisementRepository
}
type CreateAdInput struct {
	Title       string  `json:"title" binding:"required,min=3"`
	Description string  `json:"description" binding:"required,min=10"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	City        string  `json:"city"`
	ImageURL    string  `json:"image_url"`
	CategoryID  uint    `json:"category_id"`
}

func (h *AdvertisementHandler) Create(c *gin.Context) {

	var input CreateAdInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.MustGet("user_id").(uint)

	ad := models.Advertisement{
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		City:        input.City,
		ImageURL:    input.ImageURL,
		UserID:      userID,
		CategoryID:  input.CategoryID,
	}

	err := h.Repo.Create(&ad)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to create ad",
		})
		return
	}

	client := resty.New()

	_, notifyErr := client.R().
		SetBody(map[string]interface{}{
			"title":   "New Advertisement",
			"message": "New advertisement created",
		}).
		Post("http://notification-service:8082/notify")

	if notifyErr != nil {
		log.Println("failed to send notification:", notifyErr)
	}

	c.JSON(201, gin.H{
		"ad": ad,
	})
}

func (h *AdvertisementHandler) GetAll(c *gin.Context) {

	ads, err := h.Repo.GetAll()

	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to fetch ads",
		})
		return
	}

	c.JSON(200, gin.H{
		"ads": ads,
	})
}

func (h *AdvertisementHandler) GetByID(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	ad, err := h.Repo.GetByID(uint(id))

	if err != nil {
		c.JSON(404, gin.H{
			"error": "ad not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"ad": ad,
	})
}
func (h *AdvertisementHandler) Update(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	ad, err := h.Repo.GetByID(uint(id))

	if err != nil {
		c.JSON(404, gin.H{
			"error": "ad not found",
		})
		return
	}

	userID := c.MustGet("user_id").(uint)

	if ad.UserID != userID {
		c.JSON(403, gin.H{
			"error": "forbidden",
		})
		return
	}

	var input CreateAdInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ad.Title = input.Title
	ad.Description = input.Description
	ad.Price = input.Price
	ad.City = input.City
	ad.ImageURL = input.ImageURL
	ad.CategoryID = input.CategoryID

	err = h.Repo.Update(ad)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to update ad",
		})
		return
	}

	c.JSON(200, gin.H{
		"ad": ad,
	})
}
func (h *AdvertisementHandler) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	ad, err := h.Repo.GetByID(uint(id))

	if err != nil {
		c.JSON(404, gin.H{
			"error": "ad not found",
		})
		return
	}

	userID := c.MustGet("user_id").(uint)

	if ad.UserID != userID {
		c.JSON(403, gin.H{
			"error": "forbidden",
		})
		return
	}

	err = h.Repo.Delete(uint(id))

	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to delete ad",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "ad deleted",
	})
}
func (h *AdvertisementHandler) Search(c *gin.Context) {

	query := c.Query("query")

	ads, err := h.Repo.Search(query)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "search failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"ads": ads,
	})
}
func (h *AdvertisementHandler) GetMyAds(c *gin.Context) {
	userID := c.GetUint("user_id")

	ads, err := h.Repo.GetByUserID(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get ads",
		})
		return
	}

	c.JSON(http.StatusOK, ads)
}
