package user

type IUserCreateService interface {
	Handle(req UserCreateRequest) UserCreateResponse
}
