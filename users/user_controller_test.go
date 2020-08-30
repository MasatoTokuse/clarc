package users

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mtoku/di/infrastructure"
)

func TestNewUserController(t *testing.T) {
	_, err := InitializeUserController(infrastructure.NewTestMySQLConnectionString(), context.TODO())
	if err != nil {
		t.Error(err)
	}
}
func TestUserController_Create(t *testing.T) {
	controller, err := InitializeUserController(infrastructure.NewTestMySQLConnectionString(), context.TODO())
	if err != nil {
		t.Error(err)
	}

	// テスト用のリクエストとレスポンスを作成
	apiUserCreateRequest := &APIUserCreateRequest{
		UserID:   "UserController_Create.UserID",
		Password: "UserController_Create.Password",
		Nickname: "UserController_Create.Nickname",
	}
	b, err := apiUserCreateRequest.MarshalJSON()
	if err != nil {
		t.Error(err)
	}

	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(b))
	res := httptest.NewRecorder()

	controller.Create(res, req)

	// レスポンスのステータスコードのテスト
	if res.Code != http.StatusOK {
		t.Errorf("res.Code is expected : %d", res.Code)
	}

	// レスポンスのボディのテスト
	userCreateResult := &UserCreateResult{}
	err = userCreateResult.UnmarshalJSON(res.Body.Bytes())
	if err != nil {
		t.Error(err)
	}
	if userCreateResult.Status != "0" {
		t.Error("Status must be 0")
	}
	if userCreateResult.Message != "Success" {
		t.Error("Message must be Success")
	}
}
