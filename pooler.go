package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	nb_workers = 4
	workers    []*mongo.Client
)

// nb_workers : number of connection with the database
// workers : all the connected workers

func connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// uri := os.Getenv("MONGODB_URI")

	/*
		for w := 0; w < nb_workers; w++ {
			worker, err := mongo.Connect(options.Client().ApplyURI(uri))
			if err != nil {
				panic(err)
			}
			workers = append(workers, worker)
		}
	*/
}

func RemoveWorker(s []*mongo.Client, index int) []*mongo.Client {
	return append(s[:index], s[index+1:]...)
}

func close() {
	for index, worker := range workers {
		worker.Disconnect(context.Background())
		workers = RemoveWorker(workers, index)
	}
}
