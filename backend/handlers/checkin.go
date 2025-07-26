package handlers

import (
	"context"
	"fitai-backend/services"
	"github.com/gofiber/fiber/v2"
	"time"
)

func MarkCheckIn(c *fiber.Ctx) error {
	uid := c.Locals("uid").(string)
	today := time.Now().Format("2006-01-02")

	_, err := services.FirestoreClient.Collection("users").Doc(uid).
		Collection("checkins").Doc(today).
		Set(context.Background(), map[string]interface{}{
			"completed": true,
			"timestamp": time.Now(),
		})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to mark check-in"})
	}

	return c.JSON(fiber.Map{"message": "Check-in recorded"})
}
