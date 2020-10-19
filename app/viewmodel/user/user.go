package user

import "github.com/mtoku/clarc/app/viewmodel/core"

type UserCreateAPIRequest struct {
	UserID   string `json:"UserID"`
	Password string `json:"Password"`
	Name     string `json:"Name"`
}

type UserCreateAPIResult struct {
	core.APIResult
}
