package main

import (
	"rest-in-go/initializers"
	"rest-in-go/routes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Allow requests from the frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "Patch", "DELETE", "OPTIONS"}, // Allowed HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Cache-Control", "Cache", "cache", "pragma", "Expires"}, // Allowed headers
		ExposeHeaders:    []string{"Content-Length"}, // Headers exposed to the browser
		AllowCredentials: true,                      // Allow cookies
		MaxAge:           12 * time.Hour,            // Cache preflight requests
	}))

	routes.SetupRoutes(r)

	r.Run(":8080")
}
