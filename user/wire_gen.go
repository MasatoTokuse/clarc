// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package users

import (
	"context"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func InitializeUserRepository(constr MySQLConnectionString, ctx context.Context) (UserRepository, error) {
	db, err := NewMySQLDB(constr)
	if err != nil {
		return UserRepository{}, err
	}
	userRepository := NewUserRepository(ctx, db)
	return userRepository, nil
}
