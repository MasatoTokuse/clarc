//+build wireinject

package user

import (
	"context"

	"github.com/google/wire"
	"github.com/mtoku/di/app/com/domain/app/user"
	"github.com/mtoku/di/app/infrastructure"
)

func InitializeUserController(constr infrastructure.DBConnectionString, ctx context.Context) (UserController, error) {
	wire.Build(NewUserController, user.InitializeUserCreateService)
	return UserController{}, nil
}
