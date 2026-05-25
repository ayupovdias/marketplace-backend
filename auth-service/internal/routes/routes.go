package routes

import (
	"auth-service/internal/handlers"
	"auth-service/internal/middleware"
	"auth-service/internal/repository"
	"auth-service/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {

	userRepo := repository.NewUserRepository(db)

	authService := service.NewAuthService(userRepo)

	authHandler := handlers.AuthHandler{
		AuthService: authService,
		UserRepo:    userRepo,
	}

	r.GET("/health", handlers.HealthCheck)

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)

		auth.GET(
			"/profile",
			middleware.JWTAuthMiddleware(),
			authHandler.Profile,
		)
	}
}
