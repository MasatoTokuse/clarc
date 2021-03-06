package user

import (
	"context"
	"testing"

	user_domain_model "github.com/mtoku/clarc/app/domain/models/user"
	"github.com/mtoku/clarc/app/infrastructure"
)

func TestSaveFindRemoveUser(t *testing.T) {

	constr := infrastructure.NewTestMySQLConnectionString()
	repo, err := InitializeUserRepository(constr, context.TODO())
	if err != nil {
		t.Error(err)
	}
	defer repo.CloseDB()

	expected, _ := user_domain_model.NewUser(0, "mstmst11", "pass", "tks")

	result, err := repo.Save(expected)
	if err != nil {
		t.Fatal(err)
	}
	if result.ID == 0 {
		t.Fatal("user.ID must not be Zero after save")
	}

	user, err := repo.FindBy(expected.UserID.Value(), expected.Password.Value(), expected.Name.Value())
	if err != nil {
		t.Fatal(err)
	}

	repo.Remove(user)

	user, err = repo.FindBy(expected.UserID.Value(), expected.Password.Value(), expected.Name.Value())
	if err == nil {
		t.Fatal(err)
	}
}
