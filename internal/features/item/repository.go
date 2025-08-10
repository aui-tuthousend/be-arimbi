package item

import (
	"encoding/json"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemRepository interface {
	GetAll() (*[]ItemResponse, error)
	CreateItem(item Item) (*Item, error)
	UpdateItem(item Item) (*Item, error)
	GetByUuid(uuid uuid.UUID) (*Item, error)
	GetAllItemWithDetails() (*[]ItemResponseWithDetails, error)
}

type ItemRepositoryImpl struct {
	DB *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &ItemRepositoryImpl{DB: db}
}

func (ir *ItemRepositoryImpl) GetAll() (*[]ItemResponse, error) {
	var jsonData *string

	query := `
		SELECT json_agg(
			json_build_object(
				'uuid', i.uuid,
				'name', i.name,
				'description', i.description,
				'price', i.price,
				'image', i.image
			)
		)
		FROM items i
		WHERE i.deleted_at IS NULL
	`

	if err := ir.DB.Raw(query).Scan(&jsonData).Error; err != nil {
		return nil, err
	}
	items := []ItemResponse{}
	if jsonData == nil {
		return nil, nil
	}
	if err := json.Unmarshal([]byte(*jsonData), &items); err != nil {
		return nil, err
	}
	return &items, nil
}

func (ir *ItemRepositoryImpl) CreateItem(item Item) (*Item, error) {
	result := ir.DB.Create(&item)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &item, nil
}

func (ir *ItemRepositoryImpl) UpdateItem(item Item) (*Item, error) {
	result := ir.DB.Save(&item)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &item, nil
}

func (ir *ItemRepositoryImpl) GetByUuid(uuid uuid.UUID) (*Item, error) {
	var item Item
	query := `SELECT * FROM items WHERE uuid = ? AND deleted_at IS NULL`
	result := ir.DB.Raw(query, uuid).Scan(&item)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

func (ir *ItemRepositoryImpl) GetAllItemWithDetails() (*[]ItemResponseWithDetails, error) {
	var jsonData *string

	query := `
		SELECT json_agg(
			json_build_object(
				'uuid', i.uuid,
				'name', i.name,
				'description', i.description,
				'price', i.price,
				'image', i.image,
				'detail_items', (
					SELECT json_agg(
						json_build_object(
							'uuid', d.uuid,
							'name', d.name,
							'description', d.description,
							'variant', d.variant,
							'stocks', d.stocks,
							'image', d.image
						)
					)
					FROM detail_items d
					WHERE d.item_uuid = i.uuid
				)
			)
		)
		FROM items i
	`

	if err := ir.DB.Raw(query).Scan(&jsonData).Error; err != nil {
		return nil, err
	}
	response := []ItemResponseWithDetails{}
	if jsonData == nil {
		return &response, nil
	}
	if err := json.Unmarshal([]byte(*jsonData), &response); err != nil {
		return nil, err
	}
	return &response, nil
}
