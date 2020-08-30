package models

import "testing"

func Test(t *testing.T) {
	user := User{
		UserID: "1234567",
	}

	if err := user.Valid(); err == nil {
		t.Error("UserID is 7 length then, must be error")
	}

	user = User{}

	if err := user.Valid(); err == nil {
		t.Error("UserID is 0 length then, must be error")
	} else {
		if err.Error() != "userID must be more than 8 length" {
			t.Error("validation error message is wrong")
		}
	}

	user = User{
		UserID: "12345678",
	}

	if err := user.Valid(); err != nil {
		t.Error("UserID is 8 length then, must not be error")
	}
}
