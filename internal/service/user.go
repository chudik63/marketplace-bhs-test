package service

type SignInInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) SignUp(input SignInInput) {

}
