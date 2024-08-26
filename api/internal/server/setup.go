package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/infinitemax/bookAgain/internal/books"
	"github.com/infinitemax/bookAgain/internal/health"
	"github.com/infinitemax/bookAgain/internal/users"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

type Server struct {
	pool *pgxpool.Pool
}

func (s *Server) StartServer(pool *pgxpool.Pool) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	s.pool = pool

	s.SetupHandlers(r)

	err := http.ListenAndServe(":1234", r)
	if err != nil {
		fmt.Println("Fuck, the server isn't working!")
	}

	fmt.Println("listening on 1234")

}

func (s *Server) SetupHandlers(r chi.Router) {

	healthHandler := health.NewHandler()

	// users
	userDb := users.NewDb(s.pool)
	userService := users.NewService(userDb)
	userHandler := users.NewHandler(userService)

	// books
	booksDb := books.NewDb(s.pool)
	booksService := books.NewService(booksDb)
	booksHandler := books.NewHandler(booksService)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/health", func(r chi.Router) {
			r.Get("/", healthHandler.HealthCheck)
		})
		r.Route("/users", func(r chi.Router) {
			r.Post("/", userHandler.CreateUser)
			r.Get("/", userHandler.GetUsers)
		})
		r.Route("/books", func(r chi.Router) {
			r.Post("/", booksHandler.NewBook)

		})
	})

}
