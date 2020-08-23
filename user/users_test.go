package users

import (
	"context"
	"fmt"
	"testing"
)

func TestInsertUser(t *testing.T) {

	constr := NewTestMySQLConnectionString()
	repo, err := InitializeUserRepository(constr, context.TODO())
	if err != nil {
		t.Error(err)
	}
	user, err := repo.Insert()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(user)
}
