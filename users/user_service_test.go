package users

import (
	"context"
	"testing"
)

func TestNewUserService(t *testing.T) {
	_, err := InitializeUserService(NewTestMySQLConnectionString(), context.TODO())
	if err != nil {
		t.Error(err)
	}
}
