package user

// cantidad de intentos de creacion de usuarios
var attemptsCreateUsers int

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetUsers() ([]User, error) {
	return s.repo.FindAll()
}

func (s *Service) GetUserByID(id int) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *Service) CreateUser(user *User) error {
	attemptsCreateUsers++
	if err := user.Validate(); err != nil {
		return err
	}
	return s.repo.Create(user)
}

func (s *Service) UpdateUser(user *User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	return s.repo.Update(user)
}

func (s *Service) DeleteUser(id int) error {
	return s.repo.Delete(id)
}

func (s *Service) OptOutUser(id int) error {
	return s.repo.OptOut(id)
}
