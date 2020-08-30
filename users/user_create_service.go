package users

import (
	"fmt"

	"github.com/mtoku/di/models"
)

type UserCreateRequest struct {
	UserID   string
	Password string
	Nickname string
}

func (req UserCreateRequest) Valid() error {
	// IsEmpty
	if req.UserID == "" {
		return fmt.Errorf("UserID is Empty. Please enter UserID.")
	}
	if req.Password == "" {
		return fmt.Errorf("Password is Empty. Please enter Password.")
	}
	if req.Nickname == "" {
		return fmt.Errorf("Nickname is Empty. Please enter Nickname.")
	}
	return nil
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

	// Validation
	if err := req.Valid(); err != nil {
		return UserCreateResponse{
			Error: err,
		}
	}
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
