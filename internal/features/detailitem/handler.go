package detailitem

import (
	"be-arimbi/utils"
	"fmt"
	"log"
	"path/filepath"
	"strconv"

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
		stocks, err := strconv.Atoi(c.FormValue("stock"))
		if err != nil {
			return c.Status(400).JSON(utils.ErrorResponse[error](400, "stocks is required"))
		}
		req := DetailItemRequest{
			Name:        c.FormValue("name"),
			Description: c.FormValue("description"),
			Variant:     c.FormValue("variant"),
			Stocks:      stocks,
			ItemUuid:    c.FormValue("item_uuid"),
		}

		if req.ItemUuid == "" {
			return c.Status(400).JSON(utils.ErrorResponse[error](400, "item_uuid is required"))
		}

		file, err := c.FormFile("image")
		if err != nil {
			return c.Status(400).JSON(utils.ErrorResponse[error](400, "image is required"))
		}

		uploadDir := "./uploads"
		filename := fmt.Sprintf("detail-%s-%s", req.Name, file.Filename)
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
		stocks, err := strconv.Atoi(c.FormValue("stock"))
		if err != nil {
			return c.Status(400).JSON(utils.ErrorResponse[error](400, "stocks is required"))
		}
		req := DetailItemUpdateRequest{
			Uuid:        c.FormValue("uuid"),
			Name:        c.FormValue("name"),
			Description: c.FormValue("description"),
			Variant:     c.FormValue("variant"),
			Stocks:      stocks,
		}

		if req.Uuid == "" {
			return c.Status(400).JSON(utils.ErrorResponse[error](400, "uuid is required"))
		}

		newPath := ""
		file, err := c.FormFile("image")
		if file == nil {
			log.Println("image is nil")
		} else {
			uploadDir := "./uploads"
			filename := fmt.Sprintf("detail-%s-%s", req.Name, file.Filename)
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
	group := api.Group("/product-detail")
	group.Get("/:uuid", Handler.GetByUuid())
	group.Post("/", Handler.CreateDetailItem(), utils.JWTProtected())
	group.Put("/", Handler.UpdateDetailItem(), utils.JWTProtected())
}