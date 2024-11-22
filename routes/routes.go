package routes

import (
	"rest-in-go/controllers"
	"rest-in-go/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	api := r.Group("/api")

	userController := controllers.UserController{}

	api.POST("/register", userController.RegisterUser)
	api.POST("/login", userController.LoginUser)
  
	// posts
	api.GET("/posts", controllers.PostIndex)
  
	// tags 
	api.POST("/search/tags", controllers.SearchByTag)
  
	// comments
	api.GET("/posts/:id/comments", controllers.PostComments)

	protected := api.Group("/")
	
	protected.Use(middlewares.RequireAuth())
	{
		protected.GET("/profile", userController.GetProfile)
		protected.POST("/posts", controllers.PostCreate)
		protected.PATCH("/posts/:id", controllers.PostUpdate)
		protected.POST("/comments", controllers.PostComment)
	}
}
