package user

import (
	"testing"

	user_domain_model "github.com/mtoku/clarc/app/domain/models/user"
)

func TestInMemoryUserRepository(t *testing.T) {
	repo, err := InitializeInMemoryUserRepository()
	if err != nil {
		t.Error(err)
	}

	user1 := user_domain_model.User{
		UserID:   "TestInMemoryUserRepository.UserID=1",
		Password: "TestInMemoryUserRepository.Password=1",
		Name:     "TestInMemoryUserRepository.Name=1",
	}
	_, err = repo.Save(user1)
	if err != nil {
		t.Error(err)
	}
	user2 := user_domain_model.User{
		UserID:   "TestInMemoryUserRepository.UserID=2",
		Password: "TestInMemoryUserRepository.Password=2",
		Name:     "TestInMemoryUserRepository.Name=2",
	}
	_, err = repo.Save(user2)
	if err != nil {
		t.Error(err)
	}

	found, err := repo.FindBy(user1.UserID, user1.Password, user1.Name)
	if err != nil {
		t.Error(err)
	}
	if found.ID != 1 {
		t.Error("User.ID must be 1")
	}

	removed, err := repo.Remove(*found)
	if err != nil {
		t.Error(err)
	}
	if removed.ID != 1 {
		t.Error("User.ID must be 1")
	}

	removed, err = repo.FindBy(user1.UserID, user1.Password, user1.Name)
	if err == nil {
		t.Error("Not removed User.ID = 1")
	}
}
