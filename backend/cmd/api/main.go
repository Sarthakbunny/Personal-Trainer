package main

import (
	"fitai-backend/routes"
	"fitai-backend/services"
	"github.com/gofiber/fiber/v2"
	// _ "fitai-backend/docs"
	// "github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"log"
)

// @title FitAI API
// @version 1.0
// @description Backend for Personalized Diet & Workout Suggestions
// @host localhost:8080
// @BasePath /api

func main() {
	app := fiber.New()

	// Swagger documentation route
	// app.Get("/swagger/*", swagger.HandlerDefault)

	//Register all routes
	routes.RegisterRoutes(app)
	// Initialize Firebase services
	services.InitFirebase()

	app.Listen(":8080")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
