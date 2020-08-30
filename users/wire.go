//+build wireinject

package users

import (
	"context"

	"github.com/google/wire"
	"github.com/mtoku/di/infrastructure"
)

func InitializeUserRepository(constr infrastructure.DBConnectionString, ctx context.Context) (IUserRepository, error) {
	wire.Build(NewUserRepository, infrastructure.NewDB)
	return UserRepository{}, nil
}

func InitializeUserCreateService(constr infrastructure.DBConnectionString, ctx context.Context) (IUserCreateService, error) {
	wire.Build(NewUserCreateService, InitializeUserRepository)
	return UserCreateService{}, nil
}
