package users

import (
	"log"
	"net/http"
)

type Handler struct {
	s *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		s: s,
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	err := h.s.db.CreateUser()

	if err != nil {
		log.Print("Damn")
	}
}
