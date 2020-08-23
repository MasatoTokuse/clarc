//+build wireinject

package users

import (
	"context"

	"github.com/google/wire"
)

func InitializeUserRepository(constr MySQLConnectionString, ctx context.Context) (UserRepository, error) {
	wire.Build(NewUserRepository, NewMySQLDB)
	return UserRepository{}, nil
}
