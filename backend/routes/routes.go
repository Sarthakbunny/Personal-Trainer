package routes

import (
	"github.com/gofiber/fiber/v2"
	"fitai-backend/handlers"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/suggestions", handlers.GetSuggestion)
}
