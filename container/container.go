package container

import (
	"be-arimbi/internal/features/auth"
	"be-arimbi/internal/features/role"
	"be-arimbi/internal/features/user"
)

type AppContainer struct {
	AuthHandler *auth.AuthHandler
	RoleHandler *role.RoleHandler
	UserHandler *user.UserHandler
}
