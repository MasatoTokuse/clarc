package user

import "testing"

func Test(t *testing.T) {

	user := User{}

	if err := user.Valid(); err == nil {
		t.Error("UserID is empty then, must be error")
	} else {
		if err.Error() != "UserID is empty. Please enter UserID" {
			t.Error("validation error message is wrong")
		}
	}

	user = User{
		UserID: "12345678",
	}

	if err := user.Valid(); err == nil {
		t.Error("Password is empty then, must be error")
	} else {
		if err.Error() != "Password is empty. Please enter Password" {
			t.Error("validation error message is wrong")
		}
	}

	user = User{
		UserID:   "12345678",
		Password: "password",
	}

	if err := user.Valid(); err == nil {
		t.Error("Nickname is empty then, must be error")
	} else {
		if err.Error() != "Nickname is empty. Please enter Nickname" {
			t.Error("validation error message is wrong")
		}
	}

	user = User{
		UserID:   "1234567",
		Password: "password",
		Nickname: "nickname",
	}

	if err := user.Valid(); err == nil {
		t.Error("UserID is 7 length then, must be error")
	}

	user = User{
		UserID:   "12345678",
		Password: "password",
		Nickname: "nickname",
	}

	if err := user.Valid(); err != nil {
		t.Error("UserID is 8 length then, must not be error")
	}
}
