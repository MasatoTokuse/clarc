//+build wireinject

package mock

import (
	"github.com/google/wire"
	"github.com/mtoku/di/users"
)

func InitializeInMemoryUserRepository() (users.IUserRepository, error) {
	wire.Build(NewInMemoryUserRepository)
	return &InMemoryUserRepository{}, nil
}

func InitializeMockUserCreateService() (users.IUserCreateService, error) {
	wire.Build(users.NewUserCreateService, InitializeInMemoryUserRepository)
	return users.UserCreateService{}, nil
}
