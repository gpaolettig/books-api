package main

import (
	"books-api/internal/adapters/https"
	"books-api/internal/app"
	"books-api/internal/core/book"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
		panic(err)
	}
	host := os.Getenv("DB_HOST") // FIX: get the values from AWS Secrets Manager with SDK
	port := os.Getenv("DB_PORT")
	usr := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, usr, password, dbName, port)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect db %v", err))
	}
	if err := dbConn.AutoMigrate(&book.Book{}); err != nil {
		panic(fmt.Errorf("failed to migrate db: %v", err))
	}

	bookHandler := app.Init(dbConn)
	r := gin.Default()
	https.InitRoutes(r, bookHandler)
	if err := r.Run(":8080"); err != nil {
		panic(fmt.Errorf("failed to start server: %v", err))
	}

}
