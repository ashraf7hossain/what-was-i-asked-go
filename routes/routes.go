package routes

import (
	// "go/doc/comment"
	"rest-in-go/controllers"
	"rest-in-go/controllers/comment"
	"rest-in-go/controllers/post"
	"rest-in-go/controllers/user"
	"rest-in-go/controllers/vote"
	"rest-in-go/controllers/comment-vote"
	"rest-in-go/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.Use(middlewares.ErrorHandler())

	api := r.Group("/api")

	// users
	userRepository := user.NewUserRepository()
	userService := user.NewUserService(userRepository)
	UserHandler := user.NewUserHandler(userService)

	// posts
	postRepository := post.NewPostRepository()
	posetService := post.NewPostService(postRepository)
	postHandler := post.NewPostHandler(posetService)

	// comments for post
	commentRepository := comment.NewCommentRepository()
	commentService := comment.NewCommentService(commentRepository)
	commentHandler := comment.NewCommentHandler(commentService)

	// votes for post
	voteRepository := vote.NewVoteRepository()
	voteService := vote.NewVoteService(voteRepository)
	voteHandler := vote.NewVoteHandler(voteService)

	// votes for comments
	commentVoteRepository := commentvote.NewCommentVoteRepository()
	commentVoteService := commentvote.NewCommentVoteService(commentVoteRepository)
	commentVoteHandler := commentvote.NewCommentVoteHandler(commentVoteService)

	// search
	searchController := controllers.SearchController{}

	// auth
	api.POST("/register", UserHandler.CreateUser)
	api.POST("/login", UserHandler.LoginUser)

	// posts
	api.GET("/posts", postHandler.GetPosts)
	api.GET("/posts/:id", postHandler.GetPostByID)

	// search tags
	api.POST("/search/tags", searchController.SearchByTag)

	// comments
	api.GET("/posts/:id/comments", commentHandler.GetAllCommentsByPost)

	// votes
	api.GET("/posts/:id/votes", voteHandler.GetVotesByPostID)

	protected := api.Group("/")

	protected.Use(middlewares.RequireAuth(), middlewares.ExtractUserIDMiddleware())
	{
		protected.GET("/profile", UserHandler.GetProfile)
		// posts
		protected.POST("/posts", postHandler.CreatePost)
		protected.PATCH("/posts/:id", postHandler.UpdatePost)
		protected.DELETE("/posts/:id", postHandler.DeletePost)

		// comments
		protected.POST("/comments", commentHandler.PostComment)
		protected.PATCH("/comments/:comment_id", commentHandler.UpdateComment)
		protected.DELETE("/comments/:comment_id", commentHandler.DeleteComment)

		// vote post
		protected.POST("/votes", voteHandler.CreateVote)

		// vote comment
		protected.POST("/comments/votes", commentVoteHandler.CreateVote)
	}
}
