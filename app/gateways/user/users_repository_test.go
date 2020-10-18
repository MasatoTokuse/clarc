package user

import (
	"context"
	"testing"

	"github.com/mtoku/clarc/app/infrastructure"
	user_domain_model "github.com/mtoku/clarc/app/domain/models/user"
)

func TestSaveFindRemoveUser(t *testing.T) {

	constr := infrastructure.NewTestMySQLConnectionString()
	repo, err := InitializeUserRepository(constr, context.TODO())
	if err != nil {
		t.Error(err)
	}
	defer repo.CloseDB()

	expected := &user_domain_model.User{
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
