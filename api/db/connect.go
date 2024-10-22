package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func ConnectToDb() (*mongo.Client, error) {

	dbUser := os.Getenv("MONGODBUSER")
	dbPassword := os.Getenv("MONGODBPASSWORD")
	connString := fmt.Sprintf("mongodb+srv://%v:%v@cluster0.h8gi2.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0", dbUser, dbPassword)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connString).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to Database")
	}

	return client, nil
}
