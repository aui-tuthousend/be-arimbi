package user

import (
	"encoding/json"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() (*[]UserResponse, error)
	UpdateUser(user User) (*User, error)
	FindUserByUuid(uuid uuid.UUID) (*User, error)
	FindUserByEmail(email string) (*User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (ur *UserRepositoryImpl) GetAll() (*[]UserResponse, error) {
	var jsonData *string

	query := `
		SELECT json_agg(
			json_build_object(
				'uuid', u.uuid,
				'name', u.name,
				'email', u.email,
				'role', json_build_object(
					'uuid', r.uuid,
					'name', r.name
				)
			)
		)
		FROM users u
		JOIN roles r ON r.uuid = u.role_uuid
		WHERE u.deleted_at IS NULL
	`

	if err := ur.DB.Raw(query).Scan(&jsonData).Error; err != nil {
		return nil, err
	}
	users := []UserResponse{}
	if jsonData == nil {
		return nil, nil
	}
	if err := json.Unmarshal([]byte(*jsonData), &users); err != nil {
		return nil, err
	}
	return &users, nil
}


func (ur *UserRepositoryImpl) UpdateUser(user User) (*User, error) {
	result := ur.DB.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &user, nil
}

func (ur *UserRepositoryImpl) FindUserByUuid(uuid uuid.UUID) (*User, error) {
	var user User
	result := ur.DB.Raw("SELECT * FROM users WHERE uuid = ? LIMIT 1", uuid).Scan(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &user, nil
}

func (ur *UserRepositoryImpl) FindUserByEmail(email string) (*User, error) {
	var user User
	result := ur.DB.Raw("SELECT * FROM users WHERE email = ? LIMIT 1", email).Scan(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &user, nil
}
