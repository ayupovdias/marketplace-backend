package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"marketplace-service/internal/models"
	"marketplace-service/internal/repository"
	"net/http"
	"net/http/httptest"
	"testing"
)

type User struct {
	ID uint `gorm:"primaryKey"`
}

func setupTestDB() *gorm.DB {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&User{},
		&models.Advertisement{},
		&models.Category{},
		&models.Favorite{},
	)

	if err != nil {
		panic(err)
	}

	return db
}

func setupAdHandler() (*AdvertisementHandler, *gorm.DB) {

	db := setupTestDB()

	repo := repository.AdvertisementRepository{
		DB: db,
	}

	handler := AdvertisementHandler{
		Repo: &repo,
	}

	return &handler, db
}

func TestCreateAdvertisementValidation(t *testing.T) {

	gin.SetMode(gin.TestMode)

	handler, _ := setupAdHandler()

	router := gin.Default()

	router.POST("/ads", func(c *gin.Context) {
		c.Set("user_id", uint(1))
		handler.Create(c)
	})

	body := []byte(`{}`)

	req, _ := http.NewRequest(
		"POST",
		"/ads",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestCreateAdvertisementSuccess(t *testing.T) {

	gin.SetMode(gin.TestMode)

	handler, db := setupAdHandler()

	db.Create(&models.Category{
		Name: "Electronics",
	})

	router := gin.Default()

	router.POST("/ads", func(c *gin.Context) {
		c.Set("user_id", uint(1))
		handler.Create(c)
	})

	body := map[string]interface{}{
		"title":       "iPhone",
		"description": "New iPhone 15 Pro",
		"price":       1000,
		"city":        "Almaty",
		"image_url":   "https://example.com/image.jpg",
		"category_id": 1,
	}

	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest(
		"POST",
		"/ads",
		bytes.NewBuffer(jsonBody),
	)

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestGetAllAdvertisements(t *testing.T) {

	gin.SetMode(gin.TestMode)

	handler, db := setupAdHandler()

	db.Create(&models.Advertisement{
		Title:       "MacBook",
		Description: "MacBook Pro",
		Price:       2000,
		UserID:      1,
		CategoryID:  1,
	})

	router := gin.Default()

	router.GET("/ads", handler.GetAll)

	req, _ := http.NewRequest(
		"GET",
		"/ads",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetAdvertisementByID(t *testing.T) {

	gin.SetMode(gin.TestMode)

	handler, db := setupAdHandler()

	ad := models.Advertisement{
		Title:       "BMW",
		Description: "BMW M5",
		Price:       50000,
		UserID:      1,
		CategoryID:  1,
	}

	db.Create(&ad)

	router := gin.Default()

	router.GET("/ads/:id", handler.GetByID)

	req, _ := http.NewRequest(
		"GET",
		"/ads/1",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetAdvertisementNotFound(t *testing.T) {

	gin.SetMode(gin.TestMode)

	handler, _ := setupAdHandler()

	router := gin.Default()

	router.GET("/ads/:id", handler.GetByID)

	req, _ := http.NewRequest(
		"GET",
		"/ads/999",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestDeleteAdvertisementForbidden(t *testing.T) {

	gin.SetMode(gin.TestMode)

	handler, db := setupAdHandler()

	ad := models.Advertisement{
		Title:       "Laptop",
		Description: "Gaming laptop",
		Price:       1200,
		UserID:      2,
		CategoryID:  1,
	}

	db.Create(&ad)

	router := gin.Default()

	router.DELETE("/ads/:id", func(c *gin.Context) {
		c.Set("user_id", uint(1))
		handler.Delete(c)
	})

	req, _ := http.NewRequest(
		"DELETE",
		"/ads/1",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusForbidden, resp.Code)
}

func TestDeleteAdvertisementSuccess(t *testing.T) {

	gin.SetMode(gin.TestMode)

	handler, db := setupAdHandler()

	ad := models.Advertisement{
		Title:       "Camera",
		Description: "Canon Camera",
		Price:       700,
		UserID:      1,
		CategoryID:  1,
	}

	db.Create(&ad)

	router := gin.Default()

	router.DELETE("/ads/:id", func(c *gin.Context) {
		c.Set("user_id", uint(1))
		handler.Delete(c)
	})

	req, _ := http.NewRequest(
		"DELETE",
		"/ads/1",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetMyAdvertisements(t *testing.T) {

	gin.SetMode(gin.TestMode)

	handler, db := setupAdHandler()

	db.Create(&models.Advertisement{
		Title:       "PlayStation",
		Description: "PS5 Console",
		Price:       600,
		UserID:      1,
		CategoryID:  1,
	})

	router := gin.Default()

	router.GET("/ads/myads", func(c *gin.Context) {
		c.Set("user_id", uint(1))
		handler.GetMyAds(c)
	})

	req, _ := http.NewRequest(
		"GET",
		"/ads/myads",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
