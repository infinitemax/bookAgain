package main

import (
	"context"
	"fmt"
	"github.com/infinitemax/bookAgain/db/config"
	"github.com/infinitemax/bookAgain/internal/server"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connPool, err := pgxpool.NewWithConfig(context.Background(), config.Config())
	if err != nil {
		log.Fatal("Error while creating connection to the database!!")
	}

	connection, err := connPool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error while acquiring connection from the database!!")
	}
	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		log.Fatal("Could not ping database!")
	}
	fmt.Println("Connected to the database - what a relief")

	s := &server.Server{}

	s.StartServer(connPool)

}
