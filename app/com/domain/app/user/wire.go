//+build wireinject

package user

import (
	"context"

	"github.com/google/wire"
	"github.com/mtoku/di/app/gateways/user"
	"github.com/mtoku/di/app/infrastructure"
)

func InitializeUserCreateService(constr infrastructure.DBConnectionString, ctx context.Context) (IUserCreateService, error) {
	wire.Build(NewUserCreateService, user.InitializeUserRepository)
	return UserCreateService{}, nil
}
