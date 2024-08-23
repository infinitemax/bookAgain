package users

type IUserService interface {
	CreateUser() error
}

type Service struct {
	db IUserService
}

func NewService(db IUserService) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CreateUser() error {
	err := s.db.CreateUser()
	if err != nil {
		return err
	}
	return nil
}
