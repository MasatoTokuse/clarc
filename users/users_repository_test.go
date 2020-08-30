package users

import (
	"context"
	"testing"

	"github.com/mtoku/di/models"
)

func TestSaveFindRemoveUser(t *testing.T) {

	constr := NewTestMySQLConnectionString()
	repo, err := InitializeUserRepository(constr, context.TODO())
	if err != nil {
		t.Error(err)
	}
	defer repo.CloseDB()

	expected := &models.User{
		UserID:   "mst11",
		Password: "pass",
		Nickname: "tks",
	}
	result, err := repo.Save(*expected)
	if result.ID == 0 {
		t.Error("user.ID must not be Zero after save")
	}

	user, err := repo.FindBy(expected.UserID, expected.Password, expected.Nickname)
	if err != nil {
		t.Error(err)
	}

	repo.Remove(*user)

	user, err = repo.FindBy(expected.UserID, expected.Password, expected.Nickname)
	if err == nil {
		t.Error(err)
	}
}
