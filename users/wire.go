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

func InitializeUserService(constr DBConnectionString, ctx context.Context) (UserService, error) {
	wire.Build(NewUserService, InitializeUserRepository)
	return UserService{}, nil
}
