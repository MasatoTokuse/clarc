package inputdata

type APIUserCreateRequest struct {
	UserID   string `json:"UserID"`
	Password string `json:"Password"`
	Nickname string `json:"Nickname"`
}
