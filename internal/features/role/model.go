package role

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	Uuid        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	Name        string         `json:"name"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type RoleResponse struct {
	Uuid        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
}

