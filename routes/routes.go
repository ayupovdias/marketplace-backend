package routes

import (
	"marketplace/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/ads", controllers.GetAds)
	r.GET("/ads/:id", controllers.GetAd)
	r.POST("/ads", controllers.CreateAd)
	r.PUT("/ads/:id", controllers.UpdateAd)
	r.DELETE("/ads/:id", controllers.DeleteAd)

	r.GET("/ads/city/:city", controllers.GetByCity)
	r.GET("/ads/price/:min/:max", controllers.GetByPrice)
	r.GET("/ads/search", controllers.SearchAds)
	r.PATCH("/ads/:id/price", controllers.UpdatePrice)
	r.GET("/ads/latest", controllers.LatestAds)
}
