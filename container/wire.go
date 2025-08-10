//go:build wireinject
// +build wireinject

package container

import (
	"be-arimbi/internal/features/auth"
	"be-arimbi/internal/features/role"
	"be-arimbi/internal/features/user"
	"be-arimbi/internal/features/item"
	"be-arimbi/internal/features/detailitem"
	"context"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var detailItemSet = wire.NewSet(
	detailitem.NewDetailItemRepository,
	detailitem.NewDetailItemService,
	detailitem.NewDetailItemHandler,
)

var itemSet = wire.NewSet(
	item.NewItemRepository,
	item.NewItemService,
	item.NewItemHandler,
)

var roleSet = wire.NewSet(
	role.NewRoleRepository,
	role.NewRoleService,
	role.NewRoleHandler,
)

var authSet = wire.NewSet(
	auth.NewAuthRepository,
	auth.NewAuthService,
	auth.NewAuthHandler,
)

var userSet = wire.NewSet(
	user.NewUserRepository,
	user.NewUserService,
	user.NewUserHandler,
)


func InitApp(db *gorm.DB, rdb *redis.Client, ctx context.Context) *AppContainer {
	wire.Build(
		userSet,
		roleSet,
		authSet,
		itemSet,
		detailItemSet,
		wire.Struct(new(AppContainer), "*"),
	)
	return nil
}