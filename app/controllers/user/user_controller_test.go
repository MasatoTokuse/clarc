package user

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mtoku/clarc/app/infrastructure"
	"github.com/mtoku/clarc/app/testhelper"
	user_view_model "github.com/mtoku/clarc/app/viewmodel/user"
)

func TestNewUserController(t *testing.T) {
	_, err := InitializeUserController(infrastructure.NewTestMySQLConnectionString(), context.TODO())
	if err != nil {
		t.Error(err)
	}
}
func TestUserController_Create(t *testing.T) {
	db, err := infrastructure.NewDB(infrastructure.NewTestMySQLConnectionString())
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	testhelper.Cleanup(db)

	controller, err := InitializeUserController(infrastructure.NewTestMySQLConnectionString(), context.TODO())
	if err != nil {
		t.Error(err)
	}

	// テスト用のリクエストとレスポンスを作成
	apiUserCreateRequest := &user_view_model.UserCreateAPIRequest{
		UserID:   "UserController_Create.UserID",
		Password: "UserController_Create.Password",
		Name:     "UserController_Create.Name",
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
	userCreateResult := &user_view_model.UserCreateAPIResult{}
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
