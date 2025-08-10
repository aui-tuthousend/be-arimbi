package routes

import (
	"be-arimbi/container"
	"be-arimbi/internal/features/auth"
	"be-arimbi/internal/features/role"
	"be-arimbi/internal/features/user"
	"be-arimbi/internal/features/item"
	"be-arimbi/internal/features/detailitem"
	"be-arimbi/utils"

	"github.com/gofiber/fiber/v2"
)

func ProtectedRoutes(app *fiber.App, c *container.AppContainer) {
	api := app.Group("/api/admin", utils.JWTProtected())
	role.RegisterRoute(api, c.RoleHandler)
	user.RegisterRoute(api, c.UserHandler)
}

func PublicRoutes(app *fiber.App, c *container.AppContainer) {
	api := app.Group("/api")
	auth.RegisterRoute(api, c.AuthHandler)
	item.RegisterRoute(api, c.ItemHandler)
	detailitem.RegisterRoute(api, c.DetailItemHandler)
}