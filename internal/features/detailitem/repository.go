package detailitem

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DetailItemRepository interface {
	CreateDetailItem(detailItem DetailItem) (*DetailItem, error)
	UpdateDetailItem(detailItem DetailItem) (*DetailItem, error)
	GetByUuid(uuid uuid.UUID) (*DetailItem, error)
}

type DetailItemRepositoryImpl struct {
	DB *gorm.DB
}

func NewDetailItemRepository(db *gorm.DB) DetailItemRepository {
	return &DetailItemRepositoryImpl{DB: db}
}

func (dir *DetailItemRepositoryImpl) CreateDetailItem(detailItem DetailItem) (*DetailItem, error) {
	result := dir.DB.Create(&detailItem)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &detailItem, nil
}

func (dir *DetailItemRepositoryImpl) UpdateDetailItem(detailItem DetailItem) (*DetailItem, error) {
	result := dir.DB.Save(&detailItem)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &detailItem, nil
}

func (dir *DetailItemRepositoryImpl) GetByUuid(uuid uuid.UUID) (*DetailItem, error) {
	var detailItem DetailItem
	query := `SELECT * FROM detail_items WHERE uuid = ? AND deleted_at IS NULL`
	result := dir.DB.Raw(query, uuid).Scan(&detailItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &detailItem, nil
}


