package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/wignn/go-with-mongoDb/config"
	routes "github.com/wignn/go-with-mongoDb/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT") 

	if port == "" {
		port = "3000"
	}
	// Connect to MongoDB
	client := config.ConnectMongoDB()
	defer config.DisconnectMongoDB(client)

	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":" + port ))
}
