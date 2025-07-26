package handlers

import (
	"fitai-backend/services"
	"github.com/gofiber/fiber/v2"
)

// GetSuggestion godoc
// @Summary Generate personalized suggestion
// @Description Get AI-generated plan based on goal, diet, and intensity
// @Tags suggestions
// @Accept json
// @Produce json
// @Param request body SuggestionRequest true "Suggestion Input"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/suggestions [post]

type SuggestionRequest struct {
	Goal      string `json:"goal"`
	DietType  string `json:"diet_type"`
	Intensity string `json:"intensity"`
}

func GetSuggestion(c *fiber.Ctx) error {
	req := new(SuggestionRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}

	prompt := "User wants to " + req.Goal + " and follows a " + req.DietType + " diet. Give a " + req.Intensity + " intensity workout and diet plan for today."

	response, err := services.GenerateSuggestions(prompt)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "AI error"})
	}

	return c.JSON(fiber.Map{"suggestion": response})
}
