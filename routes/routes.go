package routes

import (
	  // "go/doc/comment"
	"rest-in-go/controllers"
	"rest-in-go/controllers/user"
	"rest-in-go/controllers/post"
	"rest-in-go/controllers/comment"
	"rest-in-go/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.Use(middlewares.ErrorHandler())

	api := r.Group("/api")

	userRepository := user.NewUserRepository()
	userService    := user.NewUserService(userRepository)
	UserHandler := user.NewUserHandler(userService)

	postRepository := post.NewPostRepository()
	posetService   := post.NewPostService(postRepository)
	postHandler    := post.NewPostHandler(posetService)

	commentRepository := comment.NewCommentRepository()
	commentService    := comment.NewCommentService(commentRepository)
	commentHandler    := comment.NewCommentHandler(commentService)
  
	

	// userController    := controllers.UserController{} 
	searchController  := controllers.SearchController{}
	// commentController := controllers.Comment{}

	  // auth
	api.POST("/register", UserHandler.CreateUser)
	api.POST("/login", UserHandler.LoginUser)

	  // posts
	api.GET("/posts", postHandler.GetPosts)

	  // tags
	api.POST("/search/tags", searchController.SearchByTag)

	  // comments
	api.GET("/posts/:id/comments", commentHandler.GetAllCommentsByPost)

	protected := api.Group("/")

	protected.Use(middlewares.RequireAuth(), middlewares.ExtractUserIDMiddleware())
	{
		protected.GET("/profile", UserHandler.GetProfile)
		protected.POST("/posts", postHandler.CreatePost)
		protected.PATCH("/posts/:id", postHandler.UpdatePost)
		protected.POST("/comments", commentHandler.PostComment)
	}
}
