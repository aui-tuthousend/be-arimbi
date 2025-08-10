package detailitem

import (
	"be-arimbi/internal/features/item"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DetailItem struct {
	Uuid      uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	Name      string         `json:"name"`
	Description     string         `json:"description"`
	Variant string `json:"variant"`
	Stocks int `json:"stocks"`
	Image 		string `json:"image"`
	ItemUuid uuid.UUID `json:"item_uuid"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Item item.Item `gorm:"foreignKey:ItemUuid;references:Uuid;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type DetailItemRequest struct {
	ItemUuid string `json:"item_uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Variant string `json:"variant"`
	Stocks int `json:"stocks"`
	Image string `json:"image"`
}

type DetailItemUpdateRequest struct {
	Uuid        string `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Variant string `json:"variant"`
	Stocks int `json:"stocks"`
	Image string `json:"image"`
}

