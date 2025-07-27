package handlers

import (
	"context"
	"fitai-backend/models"
	"fitai-backend/services"
	"github.com/gofiber/fiber/v2"
	"time"
)

func GetSuggestion(c *fiber.Ctx) error {
	uid := c.Locals("uid").(string)
	today := time.Now().Format("2006-01-02")

	doc, err := services.FirestoreClient.Collection("users").Doc(uid).Collection("suggestions").Doc(today).Get(context.Background())
	if err == nil && doc.Exists() {
		return c.JSON(doc.Data())
	}

	// If not exists, return a message
	return c.Status(err.Status).JSON(fiber.Map{"message": "No suggestion for today, use POST to generate one " + err.Error()})
}

func GenerateSuggestion(c *fiber.Ctx) error {
	uid := c.Locals("uid").(string)
	var req models.OnboardingData

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Build Gemini prompt
	prompt := "Create a diet and workout plan for someone with the following preferences:\n" +
		"Goal: " + req.Goal + "\n" +
		"Diet: " + req.Diet + "\n" +
		"Level: " + req.Level + "\n" +
		"Allergy: " + req.Allergy

	// Call Gemini API
	result, err := services.GenerateSuggestions(prompt)
	if err != nil {
		return c.Status(err.Status).JSON(fiber.Map{"error": "AI generation failed " + err.Error()})
	}

	// Save to Firestore
	today := time.Now().Format("2006-01-02")
	_, err = services.FirestoreClient.Collection("users").Doc(uid).
		Collection("suggestions").Doc(today).
		Set(context.Background(), map[string]interface{}{
			"date":       today,
			"suggestion": result,
		})
	if err != nil {
		return c.Status(err.Status).JSON(fiber.Map{"error": "Firestore save failed " + err.Error()})
	}

	return c.JSON(fiber.Map{"suggestion": result})
}
