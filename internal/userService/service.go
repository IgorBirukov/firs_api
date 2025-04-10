package userService

type UserService struct {
	repo UserRepository
}

func NewService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetAllUser() ([]User, error) {
	return s.repo.GetAllUser()
}

func (s *UserService) UpdateUserByID(id uint, user interface{}) (User, error) {
	return s.repo.UpdateUserByID(id, user)
}

func (s *UserService) DeleteUserByID(id uint) (res int, err error) {
	return s.repo.DeleteUserByID(id)
}
