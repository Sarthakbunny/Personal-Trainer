package routes

import (
	"fitai-backend/handlers"
	"fitai-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Protected routes with Firebase Auth
	api.Use(middleware.AuthMiddleware)

	api.Post("/onboarding", handlers.SaveOnboarding)
	api.Post("/suggestions", handlers.GenerateSuggestion)
	api.Get("/suggestions", handlers.GetSuggestions)
	api.Post("/checkin", handlers.MarkCheckIn)
	api.Get("/streak", handlers.GetStreak)
}
