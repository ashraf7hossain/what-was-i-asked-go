package routes

import (
	"rest-in-go/controllers"
	"rest-in-go/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	userController := controllers.UserController{}

	r.POST("/register", userController.RegisterUser)
	r.POST("/login", userController.LoginUser)

	r.GET("/posts", controllers.PostIndex)

	protected := r.Group("/")
	protected.Use(middlewares.RequireAuth())
	{
		protected.GET("/profile", userController.GetProfile)
		protected.POST("/posts", controllers.PostCreate)
	}
}
