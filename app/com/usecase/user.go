package usecase

import (
	"github.com/mtoku/di/app/com/usecase/inputdata"
	"github.com/mtoku/di/app/com/usecase/outputdata"
)

type IUserCreateService interface {
	Handle(req inputdata.UserCreateRequest) outputdata.UserCreateResponse
}
