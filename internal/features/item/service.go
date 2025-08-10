package item

import (
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
)

type ItemService interface {
	GetAll() (*[]ItemResponseWithDetails, error)
	GetByUuid(id string) (*ItemResponse, error)
	CreateItem(req ItemRequest, fullPath string) (*Item, error)
	UpdateItem(req ItemUpdateRequest, newPath string) (*Item, error)
}

type ItemServiceImpl struct {
	ir ItemRepository
}

func NewItemService(ir ItemRepository) ItemService {
	return &ItemServiceImpl{ir: ir}
}

func (is *ItemServiceImpl) GetAll() (*[]ItemResponseWithDetails, error) {

	items, err := is.ir.GetAllItemWithDetails()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (is *ItemServiceImpl) GetByUuid(id string) (*ItemResponse, error) {
	parsedUuid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid uuid")
	}
	
	item, err := is.ir.GetByUuid(parsedUuid)
	if err != nil {
		return nil, errors.New("item not found")
	}

	response := ItemResponse {
		Uuid: item.Uuid,
		Name: item.Name,
		Description: item.Description,
		Price: item.Price,
		Image: item.Image,
	}
	return &response, nil
}

func (is *ItemServiceImpl) CreateItem(req ItemRequest, fullPath string) (*Item, error) {

	item := Item {
		Uuid:        uuid.New(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Image:       fullPath,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return is.ir.CreateItem(item)
}

func (is *ItemServiceImpl) UpdateItem(req ItemUpdateRequest, newPath string) (*Item, error) {

	parsedUuid, err := uuid.Parse(req.Uuid)
	if err != nil {
		return nil, errors.New("invalid uuid")
	}
	
	item, err := is.ir.GetByUuid(parsedUuid)
	if err != nil {
		return nil, errors.New("item not found")
	}

	if newPath != "" {
		if err := os.Remove(item.Image); err != nil && !os.IsNotExist(err) {
			return nil, errors.New("failed to remove old image: " + err.Error())
		}
		item.Image = newPath
	}

	item.Name = req.Name
	item.Description = req.Description
	item.Price = req.Price
	item.UpdatedAt = time.Now()

	return is.ir.UpdateItem(*item)
}