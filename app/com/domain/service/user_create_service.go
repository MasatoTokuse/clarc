package service

import (
	"github.com/mtoku/di/app/com/usecase/inputdata"
	"github.com/mtoku/di/app/com/usecase/outputdata"
	"github.com/mtoku/di/app/gateways"
	"github.com/mtoku/di/app/models"
)

func NewUserCreateService(repo gateways.IUserRepository) UserCreateService {
	return UserCreateService{
		Repository: repo,
	}
}

type UserCreateService struct {
	Repository gateways.IUserRepository
}

func (service UserCreateService) Handle(req inputdata.UserCreateRequest) outputdata.UserCreateResponse {
	defer service.Repository.CloseDB()

	// Validation
	if err := req.Valid(); err != nil {
		return outputdata.UserCreateResponse{
			Error: err,
		}
	}
	user := models.User{
		UserID:   req.UserID,
		Password: req.Password,
		Nickname: req.Nickname,
	}
	if _, err := service.Repository.Save(user); err != nil {
		return outputdata.UserCreateResponse{
			Error: err,
		}
	}
	return outputdata.UserCreateResponse{}
}
