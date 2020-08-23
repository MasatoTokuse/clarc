//+build wireinject

package users

import (
	"context"

	"github.com/google/wire"
)

func InitializeUserRepository(constr DBConnectionString, ctx context.Context) (UserRepository, error) {
	wire.Build(NewUserRepository, NewDB)
	return UserRepository{}, nil
}
