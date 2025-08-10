package detailitem

import (
	"be-arimbi/internal/features/item"
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
)

type DetailItemService interface {
	GetByUuid(id string) (*item.DetailItemResponse, error)
	CreateDetailItem(req DetailItemRequest, path string) (*DetailItem, error)
	UpdateDetailItem(req DetailItemUpdateRequest, newPath string) (*DetailItem, error)
}

type DetailItemServiceImpl struct {
	dir DetailItemRepository
}

func NewDetailItemService(dir DetailItemRepository) DetailItemService {
	return &DetailItemServiceImpl{dir: dir}
}


func (dis *DetailItemServiceImpl) GetByUuid(id string) (*item.DetailItemResponse, error) {

	parsedUuid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid uuid")
	}

	detail, err := dis.dir.GetByUuid(parsedUuid)
	if err != nil {
		return nil, errors.New("detail item not found")
	}
	return &item.DetailItemResponse{
		Uuid: detail.Uuid,
		Name: detail.Name,
		Description: detail.Description,
		Variant: detail.Variant,
		Stocks: detail.Stocks,
		Image: detail.Image,
	}, nil
}


func (dis *DetailItemServiceImpl) CreateDetailItem(req DetailItemRequest, path string) (*DetailItem, error) {

	parsedUuid, err := uuid.Parse(req.ItemUuid)
	if err != nil {
		return nil, errors.New("invalid uuid")
	}

	item := DetailItem {
		Uuid: uuid.New(),
		Name: req.Name,
		Description: req.Description,
		Variant: req.Variant,
		Stocks: req.Stocks,
		Image: path,
		ItemUuid: parsedUuid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return dis.dir.CreateDetailItem(item)
}

func (dis *DetailItemServiceImpl) UpdateDetailItem(req DetailItemUpdateRequest, newPath string) (*DetailItem, error) {

	parsedUuid, err := uuid.Parse(req.Uuid)
	if err != nil {
		return nil, errors.New("invalid uuid")
	}
	
	item, err := dis.dir.GetByUuid(parsedUuid)
	if err != nil {
		return nil, errors.New("detail item not found")
	}

	if newPath != "" {
		if err := os.Remove(item.Image); err != nil && !os.IsNotExist(err) {
			return nil, errors.New("failed to remove old image: " + err.Error())
		}
		item.Image = newPath
	}

	item.Name = req.Name
	item.Description = req.Description
	item.Variant = req.Variant
	item.Stocks = req.Stocks
	item.UpdatedAt = time.Now()

	return dis.dir.UpdateDetailItem(*item)
}
