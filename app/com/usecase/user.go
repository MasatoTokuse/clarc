package usecase

import (
	"github.com/mtoku/clarc/app/com/usecase/inputdata"
	"github.com/mtoku/clarc/app/com/usecase/outputdata"
)

type IUserCreateService interface {
	Handle(req inputdata.UserCreateRequest) outputdata.UserCreateResponse
}
