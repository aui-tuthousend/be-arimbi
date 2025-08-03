package role

import (
	"be-arimbi/utils"

	"github.com/gofiber/fiber/v2"
)

type RoleHandler struct {
	rs RoleService
}

func NewRoleHandler(rs RoleService) *RoleHandler {
	return &RoleHandler{rs: rs}
}

func (rh *RoleHandler) GetAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		roles, err := rh.rs.GetAll()
		if err != nil {
			return c.Status(500).JSON(utils.ErrorResponse[error](500, err.Error()))
		}
		return c.JSON(utils.SuccessResponse(&roles))
	}
}

func RegisterRoute(api fiber.Router, Handler *RoleHandler) {
	group := api.Group("/role")
	group.Get("/", Handler.GetAll())
}