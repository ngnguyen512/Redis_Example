package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ngnguyen/shorten-url-redis/api/routes"
	//"gorm.io/gorm/logger"
)

// setupRoutes configures the Gin router with the appropriate routes.
func setupRoutes(router *gin.Engine) {
	router.GET("/:url", routes.ResolveURL)
	router.POST("/api/v1", routes.ShortenURL)
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	// Create a new Gin router instance
	router := gin.New()

	// Use custom logger (assuming you are using GORM's logger for some aspects)
	router.Use(gin.Logger())
	// You can set up a GORM logger separately if you are using GORM in your application
	// Example: db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	// Set up routes
	setupRoutes(router)

	// Retrieve the port to listen on from environment variables, default to 8080 if not set
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// Start the Gin server on the specified port
	log.Fatal(router.Run(":" + port))
}
