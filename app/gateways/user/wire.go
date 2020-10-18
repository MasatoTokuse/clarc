//+build wireinject

package user

import (
	"context"

	"github.com/google/wire"
	"github.com/mtoku/clarc/app/infrastructure"
)

func InitializeUserRepository(constr infrastructure.DBConnectionString, ctx context.Context) (UserRepository, error) {
	wire.Build(NewUserRepository, infrastructure.NewDB)
	return UserRepository{}, nil
}

func InitializeInMemoryUserRepository() (InMemoryUserRepository, error) {
	wire.Build(NewInMemoryUserRepository)
	return InMemoryUserRepository{}, nil
}
