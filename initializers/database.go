package initializers

import (
	// "fmt"
	"log"
	// "os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Global variable to store the database connection


func ConnectToDB() {
	var err error
	// password := os.Getenv("SUPABASE_PASSWORD")
	// dsn := fmt.Sprintf("host=aws-0-ap-southeast-1.pooler.supabase.com user=postgres.nnyqaziilrlawtoagpep password=%s dbname=postgres port=6543", password)
	dsn := "host=postgres user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connection established successfully")
}
