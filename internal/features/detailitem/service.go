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
	CreateDetailItem(req DetailItemRequest, path string) (*item.DetailItemResponse, error)
	UpdateDetailItem(req DetailItemUpdateRequest, newPath string) (*item.DetailItemResponse, error)
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


func (dis *DetailItemServiceImpl) CreateDetailItem(req DetailItemRequest, path string) (*item.DetailItemResponse, error) {

	parsedUuid, err := uuid.Parse(req.ItemUuid)
	if err != nil {
		return nil, errors.New("invalid uuid")
	}

	detailItem := DetailItem {
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

	createdItem, err := dis.dir.CreateDetailItem(detailItem)
	if err != nil {
		return nil, err
	}

	response := item.DetailItemResponse {
		Uuid: createdItem.Uuid,
		Name: createdItem.Name,
		Description: createdItem.Description,
		Variant: createdItem.Variant,
		Stocks: createdItem.Stocks,
		Image: createdItem.Image,
	}
	return &response, nil
}

func (dis *DetailItemServiceImpl) UpdateDetailItem(req DetailItemUpdateRequest, newPath string) (*item.DetailItemResponse, error) {

	parsedUuid, err := uuid.Parse(req.Uuid)
	if err != nil {
		return nil, errors.New("invalid uuid")
	}
	
	detailItem, err := dis.dir.GetByUuid(parsedUuid)
	if err != nil {
		return nil, errors.New("detail item not found")
	}

	if newPath != "" {
		if err := os.Remove(detailItem.Image); err != nil && !os.IsNotExist(err) {
			return nil, errors.New("failed to remove old image: " + err.Error())
		}
		detailItem.Image = newPath
	}

	detailItem.Name = req.Name
	detailItem.Description = req.Description
	detailItem.Variant = req.Variant
	detailItem.Stocks = req.Stocks
	detailItem.UpdatedAt = time.Now()

	updatedItem, err := dis.dir.UpdateDetailItem(*detailItem)
	if err != nil {
		return nil, err
	}

	response := item.DetailItemResponse {
		Uuid: updatedItem.Uuid,
		Name: updatedItem.Name,
		Description: updatedItem.Description,
		Variant: updatedItem.Variant,
		Stocks: updatedItem.Stocks,
		Image: updatedItem.Image,
	}
	return &response, nil
}
