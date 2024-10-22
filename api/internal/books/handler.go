package books

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type Handler struct {
	s      *Service
	client *mongo.Client
}

func NewHandler(r chi.Router, s *Service, client *mongo.Client) *Handler {
	h := &Handler{
		s:      s,
		client: client,
	}
	h.SetupRoutes(r)
	return h
}

func (h *Handler) SetupRoutes(r chi.Router) {

	r.Route("/books", func(r chi.Router) {
		r.Get("/", h.GetBooks)
		r.Put("/", h.NewBook)
	})

}

type Test struct {
	Test string
}

func (h *Handler) NewBook(w http.ResponseWriter, r *http.Request) {

	fmt.Println("hitting the add books route")

	testDoc := Test{
		Test: "hello",
	}

	result, err := h.client.Database("db").Collection("books").InsertOne(context.TODO(), testDoc)
	if err != nil {
		fmt.Println("error with Mongo client: ", err)
	}
	fmt.Println(result)
}

func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("look at me getting all the books!!!"))
}
