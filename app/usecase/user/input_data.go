package user

import "fmt"

type UserCreateRequest struct {
	UserID   string
	Password string
	Nickname string
}

func (req UserCreateRequest) Valid() error {
	// IsEmpty
	if req.UserID == "" {
		return fmt.Errorf("UserID is Empty. Please enter UserID.")
	}
	if req.Password == "" {
		return fmt.Errorf("Password is Empty. Please enter Password.")
	}
	if req.Nickname == "" {
		return fmt.Errorf("Nickname is Empty. Please enter Nickname.")
	}
	return nil
}
