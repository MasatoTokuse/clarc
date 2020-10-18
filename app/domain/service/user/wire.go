//+build wireinject

package user

import (
	"context"

	"github.com/google/wire"
	user_gateways "github.com/mtoku/clarc/app/gateways/user"
	"github.com/mtoku/clarc/app/infrastructure"
)

func InitializeUserCreateService(constr infrastructure.DBConnectionString, ctx context.Context) (UserCreateService, error) {
	wire.Build(UserCreateServiceSet)
	return UserCreateService{}, nil
}

var UserCreateServiceSet = wire.NewSet(NewUserCreateService,
	wire.Bind(new(user_gateways.IUserRepository), new(user_gateways.UserRepository)),
	user_gateways.InitializeUserRepository)
