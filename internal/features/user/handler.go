package user

import (
	"be-arimbi/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	us UserService
}

func NewUserHandler(us UserService) *UserHandler {
	return &UserHandler{us: us}
}

func (rh *UserHandler) GetAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := rh.us.GetAll()
		if err != nil {
			return c.Status(500).JSON(utils.ErrorResponse[error](500, err.Error()))
		}
		return c.JSON(utils.SuccessResponse(&users))
	}
}

func (rh *UserHandler) UpdateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req UserUpdateRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(utils.ErrorResponse[error](400, "Invalid request"))
		}
		user, err := rh.us.UpdateUser(req)
		if err != nil {
			return c.Status(500).JSON(utils.ErrorResponse[error](500, err.Error()))
		}
		return c.JSON(utils.SuccessResponse(&user))
	}
}

func RegisterRoute(api fiber.Router, Handler *UserHandler) {
	group := api.Group("/user")
	group.Get("/", Handler.GetAll())
	group.Put("/", Handler.UpdateUser())
}
