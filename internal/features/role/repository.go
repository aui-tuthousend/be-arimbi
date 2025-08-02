package role

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleRepository interface {
	GetAll() ([]RoleResponse, error)
	GetByUuid(uuid uuid.UUID) (*RoleResponse, error)
}

type RoleRepositoryImpl struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{DB: db}
}

func (rr *RoleRepositoryImpl) GetAll() ([]RoleResponse, error) {
	var roles []RoleResponse
	query := `SELECT uuid, name FROM roles WHERE deleted_at IS NULL`
	result := rr.DB.Raw(query).Scan(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}

func (rr *RoleRepositoryImpl) GetByUuid(uuid uuid.UUID) (*RoleResponse, error) {
	var role RoleResponse
	query := `SELECT uuid, name FROM roles WHERE uuid = ? AND deleted_at IS NULL`
	result := rr.DB.Raw(query, uuid).Scan(&role)
	if result.Error != nil {
		return nil, result.Error
	}
	return &role, nil
}

