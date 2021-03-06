package user

import (
	"context"
	"testing"

	"github.com/mtoku/clarc/app/infrastructure"
	"github.com/mtoku/clarc/app/testhelper"
	user_usecase "github.com/mtoku/clarc/app/usecase/user"
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

	req := user_usecase.UserCreateRequest{
		UserID:   "requestCreateUser.userID",
		Password: "requestCreateUser.password",
		Name:     "requestCreateUser.name",
	}

	res := iUserCreateService.Handle(req)
	if res.Error != nil {
		t.Error(res.Error)
	}
}
