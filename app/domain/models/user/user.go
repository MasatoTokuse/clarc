package user

import (
	"fmt"
)

type User struct {
	ID       int
	UserID   string
	Password string
	Nickname string
}

func (u User) Valid() error {
	if u.UserID == "" {
		return fmt.Errorf("UserID is empty. Please enter UserID")
	}
	if u.Password == "" {
		return fmt.Errorf("Password is empty. Please enter Password")
	}
	if u.Nickname == "" {
		return fmt.Errorf("Nickname is empty. Please enter Nickname")
	}
	if len(u.UserID) < 8 {
		return fmt.Errorf("userID must be more than 8 length")
	}
	return nil
}
