package user

import (
	"fmt"
)

func NewUser(id int, userID, password, name string) (*User, error) {

	user := &User{
		ID:       id,
		UserID:   UserID(userID),
		Password: Password(password),
		Name:     Name(name),
	}

	if err := user.Valid(); err != nil {
		return &User{}, err
	}

	return user, nil

}

type User struct {
	ID       int
	UserID   UserID
	Password Password
	Name     Name
}

func (u User) Valid() error {
	if err := u.UserID.Valid(); err != nil {
		return err
	}
	if err := u.Password.Valid(); err != nil {
		return err
	}
	if err := u.Name.Valid(); err != nil {
		return err
	}
	return nil
}

type UserID string

func (uid UserID) Valid() error {
	if uid == "" {
		return fmt.Errorf("UserID is empty. Please enter UserID")
	}
	if len(uid) < 8 {
		return fmt.Errorf("userID must be more than 8 length")
	}
	return nil
}

func (uid UserID) Value() string {
	return string(uid)
}

type Password string

func (password Password) Valid() error {
	if password == "" {
		return fmt.Errorf("Password is empty. Please enter Password")
	}
	return nil
}

func (password Password) Value() string {
	return string(password)
}

type Name string

func (name Name) Valid() error {
	if name == "" {
		return fmt.Errorf("Name is empty. Please enter Name")
	}
	return nil
}

func (name Name) Value() string {
	return string(name)
}
