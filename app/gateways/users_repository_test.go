package gateways

import (
	"context"
	"testing"

	"github.com/mtoku/di/app/infrastructure"
	"github.com/mtoku/di/app/models"
)

func TestSaveFindRemoveUser(t *testing.T) {

	constr := infrastructure.NewTestMySQLConnectionString()
	repo, err := InitializeUserRepository(constr, context.TODO())
	if err != nil {
		t.Error(err)
	}
	defer repo.CloseDB()

	expected := &models.User{
		UserID:   "mstmst11",
		Password: "pass",
		Nickname: "tks",
	}
	result, err := repo.Save(*expected)
	if err != nil {
		t.Fatal(err)
	}
	if result.ID == 0 {
		t.Fatal("user.ID must not be Zero after save")
	}

	user, err := repo.FindBy(expected.UserID, expected.Password, expected.Nickname)
	if err != nil {
		t.Fatal(err)
	}

	repo.Remove(*user)

	user, err = repo.FindBy(expected.UserID, expected.Password, expected.Nickname)
	if err == nil {
		t.Fatal(err)
	}
}
