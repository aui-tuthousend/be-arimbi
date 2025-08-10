package item

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	Uuid      uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	Name      string         `json:"name"`
	Description     string         `json:"description"`
	Price string `json:"price"`
	Image 		string `json:"image"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

}

type ItemResponse struct {
	Uuid        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price string `json:"price"`
	Image string `json:"image"`
}

type DetailItemResponse struct {
	Uuid        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Variant string `json:"variant"`
	Stocks int `json:"stocks"`
	Image string `json:"image"`
}

type ItemResponseWithDetails struct {
	Uuid        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price string `json:"price"`
	Image string `json:"image"`
	DetailItems []DetailItemResponse `json:"detail_items"`
}

type ItemRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price string `json:"price"`
	Image string `json:"image"`
}

type ItemUpdateRequest struct {
	Uuid        string `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price string `json:"price"`
}