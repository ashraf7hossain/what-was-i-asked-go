package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rest-in-go/initializers"
	"rest-in-go/routes"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	fmt.Println("Hello World")


  // r.GET("/posts",  controllers.PostIndex)
	// r.POST("/posts", controllers.PostCreate)
	// r.PATCH("/posts/:id", controllers.PostUpdate)

	routes.SetupRoutes(r)

	r.Run(":8080")
}
