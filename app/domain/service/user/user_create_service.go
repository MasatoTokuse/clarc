package user

import (
	user_gateways "github.com/mtoku/clarc/app/gateways/user"
	"github.com/mtoku/clarc/app/models"
	user_usecase "github.com/mtoku/clarc/app/usecase/user"
)

func NewUserCreateService(repo user_gateways.IUserRepository) UserCreateService {
	return UserCreateService{
		Repository: repo,
	}
}

type UserCreateService struct {
	Repository user_gateways.IUserRepository
}

func (service UserCreateService) Handle(req user_usecase.UserCreateRequest) user_usecase.UserCreateResponse {
	defer service.Repository.CloseDB()

	// Validation
	if err := req.Valid(); err != nil {
		return user_usecase.UserCreateResponse{
			Error: err,
		}
	}
	user := models.User{
		UserID:   req.UserID,
		Password: req.Password,
		Nickname: req.Nickname,
	}
	if _, err := service.Repository.Save(user); err != nil {
		return user_usecase.UserCreateResponse{
			Error: err,
		}
	}
	return user_usecase.UserCreateResponse{}
}
