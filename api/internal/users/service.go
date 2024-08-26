package users

import (
	"context"
)

type IUserService interface {
	GetUsers(ctx context.Context) ([]*User, error)
}

type Service struct {
	db IUserService
}

func NewService(db IUserService) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetUsers(ctx context.Context) ([]*User, error) {

	return s.db.GetUsers(ctx)
}
