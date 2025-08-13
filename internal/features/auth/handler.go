package auth

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (ah *AuthHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req LoginRequest
		log.Println(req.Email)
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		loginResponse, err := ah.authService.Login(&req)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(loginResponse)
	}
}

func RegisterRoute(api fiber.Router, Handler *AuthHandler) {
	group := api.Group("/auth")
	group.Post("/login", Handler.Login())
}
	