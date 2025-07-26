package routes

import (
	"fitai-backend/handlers"
	"fitai-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Protected routes with Firebase Auth
	protected := api.Group("/", middleware.AuthMiddleware)

	protected.Post("/onboarding", handlers.SaveOnboarding)
	protected.Get("/suggestions", handlers.GetSuggestion)
	protected.Post("/suggestions", handlers.GenerateSuggestion)
	protected.Post("/checkin", handlers.MarkCheckIn)
	protected.Get("/streak", handlers.GetStreak)
}
