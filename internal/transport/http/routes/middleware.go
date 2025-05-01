package routes

import (
	"Ads-marketplace/pkg/token"
	"github.com/gofiber/fiber/v3"
	"net/http"
)

func AuthMiddleware(c fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization token is required",
		})
	}

	tokenString = tokenString[7:]

	userId, err := token.ParseToken(tokenString)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	c.Locals("userId", userId)

	return c.Next()
}
