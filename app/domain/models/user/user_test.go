package user

import "testing"

func Test(t *testing.T) {

	_, err := NewUser(0, "", "", "")
	if err == nil {
		t.Error("UserID is empty then, must be error")
	} else {
		if err.Error() != "UserID is empty. Please enter UserID" {
			t.Error("validation error message is wrong")
		}
	}

	_, err = NewUser(0, "12345678", "", "")
	if err == nil {
		t.Error("Password is empty then, must be error")
	} else {
		if err.Error() != "Password is empty. Please enter Password" {
			t.Error("validation error message is wrong")
		}
	}

	_, err = NewUser(0, "12345678", "password", "")
	if err == nil {
		t.Error("Name is empty then, must be error")
	} else {
		if err.Error() != "Name is empty. Please enter Name" {
			t.Error("validation error message is wrong")
		}
	}

	_, err = NewUser(0, "1234567", "password", "name")
	if err == nil {
		t.Error("UserID is 7 length then, must be error")
	}

	_, err = NewUser(0, "12345678", "password", "name")
	if err != nil {
		t.Error("UserID is 8 length then, must not be error")
	}
}
