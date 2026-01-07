package utils

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func RandomStringCrypto(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

func InitClient() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")

	Client, err = mongo.Connect(options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}
}

/*
TODO: Should take in enter a schema of a database (bjson) :

	{
		"database_name": "...",
		"collections": [
			"...",
			"...",
		]
	}
*/
func CreateDatabase(databaseName string, collections []string) {
	for _, collection := range collections {
		Client.Database(databaseName).CreateCollection(context.Background(), collection)
	}
}
