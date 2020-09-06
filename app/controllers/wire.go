//+build wireinject

package controllers

import (
	"context"

	"github.com/google/wire"
	"github.com/mtoku/di/app/com/domain/service"
	"github.com/mtoku/di/app/com/usecase"
	"github.com/mtoku/di/app/infrastructure"
)

func InitializeUserController(constr infrastructure.DBConnectionString, ctx context.Context) (UserController, error) {
	wire.Build(UserControllerSet)
	return UserController{}, nil
}

var UserControllerSet = wire.NewSet(NewUserController,
	wire.Bind(new(usecase.IUserCreateService),
		new(service.UserCreateService)), service.InitializeUserCreateService)
