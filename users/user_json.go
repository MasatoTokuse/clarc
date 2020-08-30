package users

type UserCreateResult struct {
	Status  string `json:"Status"`
	Message string `json:"Message"`
}

type APIUserCreateRequest struct {
	UserID   string `json:"UserID"`
	Password string `json:"Password"`
	Nickname string `json:"Nickname"`
}
