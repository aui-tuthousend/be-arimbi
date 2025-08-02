package role

import (
	"errors"

	"github.com/google/uuid"
)

type RoleService interface {
	GetAll() ([]RoleResponse, error)
	GetByUuid(id string) (*RoleResponse, error)
}

type RoleServiceImpl struct {
	rr RoleRepository
}

func NewRoleService(rr RoleRepository) RoleService {
	return &RoleServiceImpl{rr: rr}
}

func (rs *RoleServiceImpl) GetAll() ([]RoleResponse, error) {

	roles, err := rs.rr.GetAll()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (rs *RoleServiceImpl) GetByUuid(id string) (*RoleResponse, error) {
	parsedUuid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid uuid")
	}
	role, err := rs.rr.GetByUuid(parsedUuid)
	if err != nil {
		return nil, errors.New("role not found")
	}
	return role, nil
}