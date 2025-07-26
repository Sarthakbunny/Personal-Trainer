package handlers

import (
	"context"
	"fitai-backend/services"
	"github.com/gofiber/fiber/v2"
	"time"
)

func GetStreak(c *fiber.Ctx) error {
	uid := c.Locals("uid").(string)
	checkins := services.FirestoreClient.Collection("users").Doc(uid).Collection("checkins")

	streak := 0
	today := time.Now()

	for i := 0; i < 30; i++ {
		day := today.AddDate(0, 0, -i).Format("2006-01-02")
		doc, err := checkins.Doc(day).Get(context.Background())

		if err != nil || !doc.Exists() {
			break
		}
		streak++
	}

	return c.JSON(fiber.Map{"streak": streak})
}
