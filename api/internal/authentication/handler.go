package authentication

import "net/http"

type Handler struct {
	s *Service
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {

	// check whether user exists

	// send to the service layer to do password etc

}
