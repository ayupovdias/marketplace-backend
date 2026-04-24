package routes

import (
	"marketplace/controllers"
	"marketplace/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/ads", controllers.GetAds)
	r.GET("/ads/:id", controllers.GetAd)
	
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())

	auth.POST("/ads", controllers.CreateAd)
	auth.PUT("/ads/:id", controllers.UpdateAd)
	auth.DELETE("/ads/:id", controllers.DeleteAd)

	r.GET("/ads/city/:city", controllers.GetByCity)
	r.GET("/ads/price/:min/:max", controllers.GetByPrice)
	r.GET("/ads/search", controllers.SearchAds)
	r.PATCH("/ads/:id/price", controllers.UpdatePrice)
	r.GET("/ads/latest", controllers.LatestAds)

	r.GET("/users", controllers.GetUsers)
	r.GET("/user/:id", controllers.GetUser)
	r.PUT("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	r.POST("/auth/register", controllers.Register)
	r.POST("/auth/login", controllers.Login)

	r.GET("/dashboard", middlewares.AuthMiddleware(), func(c *gin.Context) {
		user, _ := c.Get("username")
		c.JSON(200, gin.H{"message": "Welcome " + user.(string)})
	})
}
