package books

import (
	"encoding/json"
	"github.com/google/uuid"
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

func (h *Handler) NewBook(w http.ResponseWriter, r *http.Request) {

	book := &Book{}

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse("4651dca1-189e-4bee-8e12-e61304b42c9f")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.s.NewBook(r.Context(), book, id)

	_, err = w.Write([]byte("Successfully added!"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
