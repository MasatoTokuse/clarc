package users

import (
	"context"
	"fmt"
	"testing"

	"github.com/mtoku/di/models"
)

func TestRegistUser(t *testing.T) {

	constr := NewTestMySQLConnectionString()
	repo, err := InitializeUserRepository(constr, context.TODO())
	if err != nil {
		t.Error(err)
	}
	defer repo.DB.Close()
	user := models.User{
		UserID:   "mst11",
		Password: "pass",
		Nickname: "tks",
	}
	user, err = repo.Regist(user)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(user)
}

func TestFindUserBy(t *testing.T) {

	constr := NewTestMySQLConnectionString()
	repo, err := InitializeUserRepository(constr, context.TODO())
	if err != nil {
		t.Error(err)
	}
	defer repo.DB.Close()

	expected := models.User{
		UserID:   "mst11",
		Password: "",
		Nickname: "",
	}
	user, err := repo.FindBy(expected.UserID, expected.Password, expected.Nickname)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(user)
}
