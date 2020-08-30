package users

func NewUserService(repo IUserRepository) UserService {
	return UserService{
		Repository: repo,
	}
}

type UserService struct {
	Repository IUserRepository
}

// func(serv UserService)FindBy(){}
