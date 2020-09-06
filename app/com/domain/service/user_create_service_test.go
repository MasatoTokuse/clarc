package service

import (
	"context"
	"testing"

	"github.com/mtoku/di/app/com/usecase/inputdata"
	"github.com/mtoku/di/app/infrastructure"
	"github.com/mtoku/di/app/testhelper"
)

func TestNewUserCreateService(t *testing.T) {
	_, err := InitializeUserCreateService(infrastructure.NewTestMySQLConnectionString(), context.TODO())
	if err != nil {
		t.Error(err)
	}
}

func TestUserCreateService(t *testing.T) {
	db, err := infrastructure.NewDB(infrastructure.NewTestMySQLConnectionString())
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	testhelper.Cleanup(db)

	iUserCreateService, err := InitializeUserCreateService(infrastructure.NewTestMySQLConnectionString(), context.TODO())
	if err != nil {
		t.Error(err)
	}

	req := inputdata.UserCreateRequest{
		UserID:   "requestCreateUser.userID",
		Password: "requestCreateUser.password",
		Nickname: "requestCreateUser.nickname",
	}

	res := iUserCreateService.Handle(req)
	if res.Error != nil {
		t.Error(res.Error)
	}
}
