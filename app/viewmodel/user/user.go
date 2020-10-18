package user

import "github.com/mtoku/clarc/app/viewmodel/core"

type UserCreateAPIRequest struct {
	UserID   string `json:"UserID"`
	Password string `json:"Password"`
	Nickname string `json:"Nickname"`
}

type UserCreateAPIResult struct {
	core.APIResult
}
