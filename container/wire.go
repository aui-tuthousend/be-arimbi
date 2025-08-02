//go:build wireinject
// +build wireinject

package container


import (
	"be-arimbi/internal/features/auth"
	"be-arimbi/internal/features/role"
	"be-arimbi/internal/features/user"

	"github.com/google/wire"
	"gorm.io/gorm"
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


func InitApp(db *gorm.DB) *AppContainer {
	wire.Build(
		userSet,
		roleSet,
		authSet,
		wire.Struct(new(AppContainer), "*"),
	)
	return nil
}