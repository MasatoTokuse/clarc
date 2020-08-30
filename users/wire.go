//+build wireinject

package users

import (
	"context"

	"github.com/google/wire"
)

func InitializeUserRepository(constr DBConnectionString, ctx context.Context) (IUserRepository, error) {
	wire.Build(NewUserRepository, NewDB)
	return UserRepository{}, nil
}

func InitializeUserCreateService(constr DBConnectionString, ctx context.Context) (IUserCreateService, error) {
	wire.Build(NewUserCreateService, InitializeUserRepository)
	return UserCreateService{}, nil
}
