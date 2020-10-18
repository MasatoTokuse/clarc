//+build wireinject

package service

import (
	"context"

	"github.com/google/wire"
	"github.com/mtoku/clarc/app/gateways"
	"github.com/mtoku/clarc/app/infrastructure"
)

func InitializeUserCreateService(constr infrastructure.DBConnectionString, ctx context.Context) (UserCreateService, error) {
	wire.Build(UserCreateServiceSet)
	return UserCreateService{}, nil
}

var UserCreateServiceSet = wire.NewSet(NewUserCreateService,
	wire.Bind(new(gateways.IUserRepository), new(gateways.UserRepository)),
	gateways.InitializeUserRepository)
