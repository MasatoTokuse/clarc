package users

import "github.com/mtoku/di/models"

type UserCreateRequest struct {
	UserID   string
	Password string
	Nickname string
}

type UserCreateResponse struct {
	Error error
}

type IUserCreateService interface {
	Handle(req UserCreateRequest) UserCreateResponse
}

func NewUserCreateService(repo IUserRepository) IUserCreateService {
	return UserCreateService{
		Repository: repo,
	}
}

type UserCreateService struct {
	Repository IUserRepository
}

func (service UserCreateService) Handle(req UserCreateRequest) UserCreateResponse {
	defer service.Repository.CloseDB()
	user := models.User{
		UserID:   req.UserID,
		Password: req.Password,
		Nickname: req.Nickname,
	}
	if _, err := service.Repository.Save(user); err != nil {
		return UserCreateResponse{
			Error: err,
		}
	}
	return UserCreateResponse{}
}
