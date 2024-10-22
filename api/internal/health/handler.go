package health

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewHandler(r chi.Router) {
	r.Get("/health", GetHealth)
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
