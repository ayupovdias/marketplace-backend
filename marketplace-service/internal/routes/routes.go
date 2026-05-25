package routes

import (
	"marketplace-service/internal/handlers"
	"marketplace-service/internal/middleware"
	"marketplace-service/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {

	adRepo := repository.NewAdvertisementRepository(db)

	adHandler := handlers.AdvertisementHandler{
		Repo: adRepo,
	}

	categoryHandler := handlers.CategoryHandler{
		DB: db,
	}

	favoriteHandler := handlers.FavoriteHandler{
		DB: db,
	}

	r.GET("/health", handlers.HealthCheck)

	ads := r.Group("/ads")
	{
		ads.GET("/", adHandler.GetAll)
		ads.GET("/:id", adHandler.GetByID)
		ads.GET("/search", adHandler.Search)

		protected := ads.Group("/")
		protected.Use(middleware.JWTAuthMiddleware())

		protected.POST("/", adHandler.Create)
		protected.PUT("/:id", adHandler.Update)
		protected.DELETE("/:id", adHandler.Delete)

		protected.GET("/my", adHandler.GetMyAds)
	}

	categories := r.Group("/categories")
	{
		categories.GET("/", categoryHandler.GetAll)
		categories.GET("/:id", categoryHandler.GetByID)

		protected := categories.Group("/")
		protected.Use(middleware.JWTAuthMiddleware())

		protected.POST("/", categoryHandler.Create)
		protected.PUT("/:id", categoryHandler.Update)
		protected.DELETE("/:id", categoryHandler.Delete)
	}

	favorites := r.Group("/favorites")
	favorites.Use(middleware.JWTAuthMiddleware())
	{
		favorites.POST("/:ad_id", favoriteHandler.Add)
		favorites.GET("/", favoriteHandler.GetAll)

		favorites.GET("/:id", favoriteHandler.GetByID)
		favorites.DELETE("/:id", favoriteHandler.Delete)
	}
}
