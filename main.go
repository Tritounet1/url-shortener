package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Book struct {
	Title     string    `json:"title" bson:"title"`
	Author    string    `json:"author" bson:"author"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdateAt  time.Time `json:"update_at" bson:"update_at"`
}

func NewBook(title string, author string) Book {
	book := Book{}
	book.Title = title
	book.Author = author
	book.CreatedAt = time.Now()
	book.UpdateAt = time.Now()
	return book
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Select a collection
	coll := client.Database("db").Collection("books")

	// Create the document for insert
	doc := NewBook("Atonement", "Ian McEwan")

	// Insert the new document
	result, err := coll.InsertOne(context.TODO(), doc)

	// Print the ID of the document (automatically create by mongo)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

	/*
		coll := client.Database("sample_mflix").Collection("movies")
		title := "Back to the Future"

		var result bson.M
		err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).
			Decode(&result)
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No document was found with the title %s\n", title)
			return
		}
		if err != nil {
			panic(err)
		}

		jsonData, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", jsonData)
	*/
}
