package books

import (
	"context"
	"github.com/google/uuid"
)

type IBooksHandler interface {
	NewBook(ctx context.Context, book *Book, user uuid.UUID) error
}

type Service struct {
	db IBooksHandler
}

func NewService(db IBooksHandler) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) NewBook(ctx context.Context, book *Book, user uuid.UUID) error {
	return s.db.NewBook(ctx, book, user)
}
