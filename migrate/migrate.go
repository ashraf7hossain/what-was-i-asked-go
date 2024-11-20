package main

import(
	"log"
	"rest-in-go/initializers"
	"rest-in-go/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
    err := initializers.DB.AutoMigrate(&models.Post{})
    if err != nil {
        log.Fatal("Failed to migrate Post model:", err)
    }

    err = initializers.DB.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Failed to migrate User model:", err)
    }

    err = initializers.DB.AutoMigrate(&models.Comment{})
    if err != nil {
        log.Fatal("Failed to migrate Comment model:", err)
    }

    err = initializers.DB.AutoMigrate(&models.Vote{})
    if err != nil {
        log.Fatal("Failed to migrate Vote model:", err)
    }

    log.Println("Migrations completed successfully")
}
