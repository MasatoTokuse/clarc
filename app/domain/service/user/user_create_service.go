package user

import (
	user_domain_model "github.com/mtoku/clarc/app/domain/models/user"
	user_gateways "github.com/mtoku/clarc/app/gateways/user"
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
	user, err := user_domain_model.NewUser(0, req.UserID, req.Password, req.Name)
	if err != nil {
		return user_usecase.UserCreateResponse{
			Error: err,
		}
	}
	if _, err := service.Repository.Save(user); err != nil {
		return user_usecase.UserCreateResponse{
			Error: err,
		}
	}
	return user_usecase.UserCreateResponse{}
}
