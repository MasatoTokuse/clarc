//+build wireinject

package user

import (
	"context"

	"github.com/google/wire"
	"github.com/mtoku/clarc/app/domain/service/user"
	"github.com/mtoku/clarc/app/infrastructure"
	user_usecase "github.com/mtoku/clarc/app/usecase/user"
)

func InitializeUserController(constr infrastructure.DBConnectionString, ctx context.Context) (UserController, error) {
	wire.Build(UserControllerSet)
	return UserController{}, nil
}

var UserControllerSet = wire.NewSet(NewUserController,
	wire.Bind(new(user_usecase.IUserCreateService),
		new(user.UserCreateService)), user.InitializeUserCreateService)
