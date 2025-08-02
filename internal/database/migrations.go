package database

import (
	"be-arimbi/internal/features/role"
	"be-arimbi/internal/features/user"
	"be-arimbi/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func MigrateAndSeed(db *gorm.DB) error {

	var existing role.Role
	if err := db.Where("name = ?", "admin").First(&existing).Error; err == gorm.ErrRecordNotFound {
		// Seeder
		roles := []role.Role{
			{Uuid: uuid.New(), Name: "admin"},
			{Uuid: uuid.New(), Name: "user"},
		}

		for _, Role := range roles {
			db.Create(&Role)
		}

		hashedPassword, _ := utils.HashPassword("admin")
		user := user.User{
			Uuid:        uuid.New(),
			Name:        "admin",
			Email:       "admin@admin.com",
			Password:    hashedPassword,
			RoleUuid:    roles[0].Uuid,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		db.Create(&user)
	}

	return nil
}
