package users

import (
	"context"
	"testing"
)

func TestNewUserCreateService(t *testing.T) {
	_, err := InitializeUserCreateService(NewTestMySQLConnectionString(), context.TODO())
	if err != nil {
		t.Error(err)
	}
}

func TestUserCreateService(t *testing.T) {
	iUserCreateService, err := InitializeUserCreateService(NewTestMySQLConnectionString(), context.TODO())
	if err != nil {
		t.Error(err)
	}

	req := UserCreateRequest{
		UserID:   "requestCreateUser.userID",
		Password: "requestCreateUser.password",
		Nickname: "requestCreateUser.nickname",
	}

	res := iUserCreateService.Handle(req)
	if res.Error != nil {
		t.Error(res.Error)
	}
}
