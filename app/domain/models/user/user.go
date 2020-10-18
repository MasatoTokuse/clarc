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
	if len(u.UserID) < 8 {
		return fmt.Errorf("userID must be more than 8 length")
	}
	return nil
}
