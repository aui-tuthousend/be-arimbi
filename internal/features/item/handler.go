package item

import (
	"be-arimbi/utils"
	"fmt"
	"log"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

type ItemHandler struct {
	is ItemService
}

func NewItemHandler(is ItemService) *ItemHandler {
	return &ItemHandler{is: is}
}

func (ih *ItemHandler) GetAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		items, err := ih.is.GetAll()
		if err != nil {
			return c.Status(500).JSON(utils.ErrorResponse[error](500, err.Error()))
		}
		return c.JSON(utils.SuccessResponse(&items))
	}
}

func (ih *ItemHandler) GetByUuid() fiber.Handler {
	return func(c *fiber.Ctx) error {
		uuid := c.Params("uuid")
		item, err := ih.is.GetByUuid(uuid)
		if err != nil {
			return c.Status(500).JSON(utils.ErrorResponse[error](500, err.Error()))
		}
		return c.JSON(utils.SuccessResponse(&item))
	}
}

func (ih *ItemHandler) CreateItem() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req ItemRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(utils.ErrorResponse[error](400, "Invalid request"))
		}

		file, err := c.FormFile("image")
		if err != nil {
			return c.Status(400).JSON(utils.ErrorResponse[error](400, "image cover is required"))
		}

		uploadDir := "./uploads"
		filename := fmt.Sprintf("%s-%s", file.Filename, "cover")
		fullPath := filepath.Join(uploadDir, filename)

		if err := c.SaveFile(file, fullPath); err != nil {
			return c.Status(500).JSON(utils.ErrorResponse[error](500, "failed to save file"))
		}

		item, err := ih.is.CreateItem(req, fullPath)
		if err != nil {
			return c.Status(500).JSON(utils.ErrorResponse[error](500, err.Error()))
		}
		return c.JSON(utils.SuccessResponse(&item))
	}
}

func (ih *ItemHandler) UpdateItem() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req ItemUpdateRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(utils.ErrorResponse[error](400, "Invalid request"))
		}

		newPath := ""
		file, err := c.FormFile("image")
		if file == nil {
			log.Println("file is nil")
		} else {
			uploadDir := "./uploads"
			filename := fmt.Sprintf("%s-%s", file.Filename, "cover")
			newPath = filepath.Join(uploadDir, filename)

			if err := c.SaveFile(file, newPath); err != nil {
				return c.Status(500).JSON(utils.ErrorResponse[error](500, "failed to save file"))
			}
		}
		
		item, err := ih.is.UpdateItem(req, newPath)
		if err != nil {
			return c.Status(500).JSON(utils.ErrorResponse[error](500, err.Error()))
		}
		return c.JSON(utils.SuccessResponse(&item))
	}
}

func RegisterRoute(api fiber.Router, Handler *ItemHandler) {
	group := api.Group("/item")
	group.Get("/", Handler.GetAll())
	group.Get("/:uuid", Handler.GetByUuid())
	group.Post("/", Handler.CreateItem(), utils.JWTProtected())
	group.Put("/", Handler.UpdateItem(), utils.JWTProtected())
}
