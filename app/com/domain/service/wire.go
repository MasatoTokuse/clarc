//+build wireinject

package service

import (
	"context"

	"github.com/google/wire"
	"github.com/mtoku/di/app/com/usecase"
	"github.com/mtoku/di/app/gateways"
	"github.com/mtoku/di/app/infrastructure"
)

func InitializeUserCreateService(constr infrastructure.DBConnectionString, ctx context.Context) (usecase.IUserCreateService, error) {
	wire.Build(NewUserCreateService, gateways.InitializeUserRepository)
	return UserCreateService{}, nil
}
