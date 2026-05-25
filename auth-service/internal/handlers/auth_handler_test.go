package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterValidation(t *testing.T) {

	router := gin.Default()

	handler := AuthHandler{}

	router.POST("/register", handler.Register)

	body := []byte(`{
		"email":"test@test.com"
	}`)

	req, _ := http.NewRequest(
		"POST",
		"/register",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
}

func TestLoginValidation(t *testing.T) {

	router := gin.Default()

	handler := AuthHandler{}

	router.POST("/login", handler.Login)

	body := []byte(`{}`)

	req, _ := http.NewRequest(
		"POST",
		"/login",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
}
