package users

func NewUserService(repo UserRepository) UserService {
	return UserService{
		Repository: repo,
	}
}

type UserService struct {
	Repository UserRepository
}

// func(serv UserService)FindBy(){}
