package server

import (
	"github.com/gin-gonic/gin"
	"github.com/quyenphamkhac/emojition/controllers"
)

// NewRouter gin
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	api := router.Group("api")
	{
		userController := new(controllers.UserController)
		api.POST("/signup", userController.SignUp)
		api.POST("/login", userController.Login)
		productGroup := api.Group("product")
		{
			product := new(controllers.ProductController)
			productGroup.POST("/", product.Create)
		}
	}
	return router
}
