package user

import (
	"time"
	"be-arimbi/internal/features/role"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Uuid      uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"uuid"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	RoleUuid  uuid.UUID      `gorm:"type:uuid" json:"role_uuid"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Role role.Role `gorm:"foreignKey:RoleUuid;references:Uuid;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type UserLoginResponse struct {
	Uuid        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Role        string `json:"role"`
}

type UserResponse struct {
	Uuid        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Role        role.RoleResponse `json:"role"`
}

type UserRequest struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
}

type UserUpdateRequest struct {
	Uuid        string `json:"uuid"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"role"`
}