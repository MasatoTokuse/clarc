//+build wireinject

package user

import (
	"context"

	"github.com/google/wire"
	"github.com/mtoku/di/app/infrastructure"
)

func InitializeUserRepository(constr infrastructure.DBConnectionString, ctx context.Context) (IUserRepository, error) {
	wire.Build(NewUserRepository, infrastructure.NewDB)
	return UserRepository{}, nil
}

func InitializeInMemoryUserRepository() (IUserRepository, error) {
	wire.Build(NewInMemoryUserRepository)
	return &InMemoryUserRepository{}, nil
}
