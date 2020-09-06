package request

type UserCreateAPIRequest struct {
	UserID   string `json:"UserID"`
	Password string `json:"Password"`
	Nickname string `json:"Nickname"`
}
