package user

import (
	"fmt"

	user_domain_model "github.com/mtoku/clarc/app/domain/models/user"
)

func NewInMemoryUserRepository() InMemoryUserRepository {
	return InMemoryUserRepository{increment: 1}
}

type InMemoryUserRepository struct {
	increment int
	Users     []user_domain_model.User
}

func (repo *InMemoryUserRepository) Save(user user_domain_model.User) (user_domain_model.User, error) {

	if err := user.Valid(); err != nil {
		return user_domain_model.User{}, err
	}

	user.ID = repo.increment
	repo.Users = append(repo.Users, user)
	repo.increment++

	return user, nil
}

func (repo *InMemoryUserRepository) FindBy(userID, password, name string) (user_domain_model.User, error) {

	ret := user_domain_model.User{}
	err := fmt.Errorf("No users")

	for _, user := range repo.Users {
		if userID != "" && user.UserID.Value() != userID {
			continue
		}
		if password != "" && user.Password.Value() != password {
			continue
		}
		if name != "" && user.Name.Value() != name {
			continue
		}
		ret = user
		err = nil
		break
	}

	return ret, err
}

func (repo *InMemoryUserRepository) Remove(target user_domain_model.User) (user_domain_model.User, error) {

	ret := user_domain_model.User{}
	err := fmt.Errorf("No users")

	for i, user := range repo.Users {
		if target.UserID != "" && user.UserID != target.UserID {
			continue
		}
		if target.Password != "" && user.Password != target.Password {
			continue
		}
		if target.Name != "" && user.Name != target.Name {
			continue
		}

		repo.Users = append(repo.Users[0:i], repo.Users[i+1:len(repo.Users)]...)

		ret = user
		err = nil
		break
	}

	return ret, err
}

func (repo *InMemoryUserRepository) CloseDB() error {
	return nil
}
