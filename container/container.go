package container

import (
	"be-arimbi/internal/features/auth"
	"be-arimbi/internal/features/role"
	"be-arimbi/internal/features/user"
	"be-arimbi/internal/features/item"
	"be-arimbi/internal/features/detailitem"
)

type AppContainer struct {
	AuthHandler *auth.AuthHandler
	RoleHandler *role.RoleHandler
	UserHandler *user.UserHandler
	ItemHandler *item.ItemHandler
	DetailItemHandler *detailitem.DetailItemHandler
}
