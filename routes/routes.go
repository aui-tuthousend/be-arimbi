package routes

import (
	"be-arimbi/container"
	"be-arimbi/internal/features/auth"
	"be-arimbi/internal/features/role"
	"be-arimbi/internal/features/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, c *container.AppContainer) {
	api := app.Group("/api")

	role.RegisterRoute(api, c.RoleHandler)
	user.RegisterRoute(api, c.UserHandler)
	auth.RegisterRoute(api, c.AuthHandler)
}
