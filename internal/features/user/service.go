package user

import (
	"be-arimbi/utils"
	"errors"

	"github.com/google/uuid"
)

type UserService interface {
	GetAll() ([]UserResponse, error)
	UpdateUser(req UserUpdateRequest) (*User, error)
}

type UserServiceImpl struct {
	ur UserRepository
}

func NewUserService(ur UserRepository) UserService {
	return &UserServiceImpl{ur: ur}
}

func (us *UserServiceImpl) GetAll() ([]UserResponse, error) {
	return us.ur.GetAll()
}

func (us *UserServiceImpl) UpdateUser(req UserUpdateRequest) (*User, error) {

	parsedUuid, err := uuid.Parse(req.Uuid)
	if err != nil {
		return nil, errors.New("invalid uuid")
	}
	user, err := us.ur.FindUserByUuid(parsedUuid)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	user.Name = req.Name
	user.Email = req.Email

	if req.Password != "" {
		hashedPassword, _ := utils.HashPassword(req.Password)
		user.Password = hashedPassword
	}

	return us.ur.UpdateUser(*user)
}