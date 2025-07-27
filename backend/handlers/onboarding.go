package handlers

import (
	"cloud.google.com/go/firestore"
	"context"
	"fitai-backend/models"
	"fitai-backend/services"
	"github.com/gofiber/fiber/v2"
)

func SaveOnboarding(c *fiber.Ctx) error {
	uid := c.Locals("uid").(string)
	var data models.OnboardingData

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}

	_, err := services.FirestoreClient.Collection("users").Doc(uid).Set(context.Background(), map[string]interface{}{
		"onboarding": data,
	}, firestore.MergeAll) // Use MergeAll to update or create the document
	if err != nil {
		return c.Status(err.Status).JSON(fiber.Map{"error": "Failed to save onboarding " + err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Onboarding saved"})
}
