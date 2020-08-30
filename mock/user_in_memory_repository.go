package mock

import (
	"fmt"

	"github.com/mtoku/di/models"
	"github.com/mtoku/di/users"
)

func NewInMemoryUserRepository() users.IUserRepository {
	return &InMemoryUserRepository{increment: 1}
}

type InMemoryUserRepository struct {
	increment int
	Users     []models.User
}

func (repo *InMemoryUserRepository) Save(user models.User) (*models.User, error) {

	if err := user.Valid(); err != nil {
		return &models.User{}, err
	}

	user.ID = repo.increment
	repo.Users = append(repo.Users, user)
	repo.increment++

	return &user, nil
}

func (repo *InMemoryUserRepository) FindBy(userID, password, nickname string) (*models.User, error) {

	ret := models.User{}
	err := fmt.Errorf("No users")

	for _, user := range repo.Users {
		if userID != "" && user.UserID != userID {
			continue
		}
		if password != "" && user.Password != password {
			continue
		}
		if nickname != "" && user.Nickname != nickname {
			continue
		}
		ret = user
		err = nil
		break
	}

	return &ret, err
}

func (repo *InMemoryUserRepository) Remove(target models.User) (*models.User, error) {

	ret := models.User{}
	err := fmt.Errorf("No users")

	for i, user := range repo.Users {
		if target.UserID != "" && user.UserID != target.UserID {
			continue
		}
		if target.Password != "" && user.Password != target.Password {
			continue
		}
		if target.Nickname != "" && user.Nickname != target.Nickname {
			continue
		}

		repo.Users = append(repo.Users[0:i], repo.Users[i+1:len(repo.Users)]...)

		ret = user
		err = nil
		break
	}

	return &ret, err
}

func (repo *InMemoryUserRepository) CloseDB() error {
	return nil
}
