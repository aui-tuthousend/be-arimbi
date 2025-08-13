package detailitem

import (
	"be-arimbi/utils"
	"fmt"
	"log"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

type DetailItemHandler struct {
	dis DetailItemService
}

func NewDetailItemHandler(dis DetailItemService) *DetailItemHandler {
	return &DetailItemHandler{dis: dis}
}

func (ih *DetailItemHandler) GetByUuid() fiber.Handler {
	return func(c *fiber.Ctx) error {
		uuid := c.Params("uuid")
		item, err := ih.dis.GetByUuid(uuid)
		if err != nil {
			return c.Status(500).JSON(utils.ErrorResponse[error](500, err.Error()))
		}
		return c.JSON(utils.SuccessResponse(&item))
	}
}

func (ih *DetailItemHandler) CreateDetailItem() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req DetailItemRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(utils.ErrorResponse[error](400, err.Error()))
		}

		file, err := c.FormFile("image")
		if err != nil {
			return c.Status(400).JSON(utils.ErrorResponse[error](400, "image is required"))
		}

		uploadDir := "./uploads"
		filename := fmt.Sprintf("%s-%s", file.Filename, "detail")
		fullPath := filepath.Join(uploadDir, filename)

		if err := c.SaveFile(file, fullPath); err != nil {
			return c.Status(500).JSON(utils.ErrorResponse[error](500, "failed to save file"))
		}

		createdDetailItem, err := ih.dis.CreateDetailItem(req, fullPath)
		if err != nil {
			return c.Status(500).JSON(utils.ErrorResponse[error](500, err.Error()))
		}
		return c.JSON(utils.SuccessResponse(&createdDetailItem))
	}
}

func (ih *DetailItemHandler) UpdateDetailItem() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req DetailItemUpdateRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(utils.ErrorResponse[error](400, "Invalid request"))
		}

		newPath := ""
		file, err := c.FormFile("image")
		if file == nil {
			log.Println("image is nil")
		} else {
			uploadDir := "./uploads"
			filename := fmt.Sprintf("%s-%s", file.Filename, "detail")
			newPath = filepath.Join(uploadDir, filename)

			if err := c.SaveFile(file, newPath); err != nil {
				return c.Status(500).JSON(utils.ErrorResponse[error](500, "failed to save file"))
			}
		}
		
		item, err := ih.dis.UpdateDetailItem(req, newPath)
		if err != nil {
			return c.Status(500).JSON(utils.ErrorResponse[error](500, err.Error()))
		}
		return c.JSON(utils.SuccessResponse(&item))
	}
}

func RegisterRoute(api fiber.Router, Handler *DetailItemHandler) {
	group := api.Group("/detail-product")
	group.Get("/:uuid", Handler.GetByUuid())
	group.Post("/", Handler.CreateDetailItem(), utils.JWTProtected())
	group.Put("/", Handler.UpdateDetailItem(), utils.JWTProtected())
}