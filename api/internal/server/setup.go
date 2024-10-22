package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/infinitemax/bookAgain/internal/books"
	"github.com/infinitemax/bookAgain/internal/health"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Client *mongo.Client
	Deps   *dependencies
}

type dependencies struct {
	booksService *books.Service
}

func (s *Server) SetupDependencies() error {
	s.Deps = &dependencies{}

	booksService := books.NewService()
	s.Deps.booksService = booksService

	return nil
}

func (s *Server) SetupHandlers(r chi.Router) error {
	r.Use(middleware.Logger)

	err := s.SetupDependencies()
	if err != nil {
		return err
	}
	health.NewHandler(r)

	books.NewHandler(r, s.Deps.booksService, s.Client)
	return nil
}
