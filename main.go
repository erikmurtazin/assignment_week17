package main

import (
	"assignment_week17/api"
	"assignment_week17/db"
	"context"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load the environment variables")
	}
	store, err := db.NewStorage()
	if err != nil {
		log.Fatal("Failed to connect to mongo")
	}
	defer func() {
		if err := store.Client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
	s := api.NewServer(store)
	s.Run()
}
