package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/infinitemax/bookAgain/db"
	"github.com/infinitemax/bookAgain/internal/server"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func main() {

	svr := server.Server{}

	// connect to db
	client, err := db.ConnectToDb()
	if err != nil {
		log.Fatal("problem connecting to mongodb.")
	}
	defer client.Disconnect(context.Background())

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	svr.Client = client

	// create router and handlers
	r := chi.NewRouter()
	err = svr.SetupHandlers(r)
	if err != nil {
		log.Fatal("problem setting up handlers.")
	}

	// start server
	log.Fatal(http.ListenAndServe(":2811", r))
}
