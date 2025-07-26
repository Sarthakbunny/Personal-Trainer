package middleware

import (
	"context"
	"fitai-backend/services"
	"github.com/gofiber/fiber/v2"
	"strings"
)

// AuthMiddleware validates Firebase token and adds UID to context
func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
	}

	idToken := strings.Replace(authHeader, "Bearer ", "", 1)

	// Verify token
	token, err := services.FirebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	// Attach UID to context
	c.Locals("uid", token.UID)
	return c.Next()
}
