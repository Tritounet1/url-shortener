package utils

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"tidy/models"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"golang.org/x/crypto/bcrypt"
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

	// Create unique index on username in token collection (one-to-one relationship)
	CreateUniqueIndexOnToken(databaseName)
}

func CreateUniqueIndexOnToken(databaseName string) {
	collection := Client.Database(databaseName).Collection("token")

	indexModel := mongo.IndexModel{
		Keys:    map[string]interface{}{"username": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Printf("Warning: Could not create unique index on token.username: %v", err)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateJWTToken(username string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		println("err : ", err.Error())
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func SaveToken(username string, tokenString string) error {
	collection := Client.Database("db").Collection("token")

	filter := bson.D{{"username", username}}
	token := models.NewToken(tokenString, username)

	opts := options.UpdateOne().SetUpsert(true)
	update := bson.D{{"$set", token}}

	_, err := collection.UpdateOne(context.Background(), filter, update, opts)
	return err
}

func GetTokenByUsername(username string) (*models.Token, error) {
	collection := Client.Database("db").Collection("token")

	filter := bson.D{{"username", username}}
	var token models.Token

	err := collection.FindOne(context.Background(), filter).Decode(&token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func DeleteTokenByUsername(username string) error {
	collection := Client.Database("db").Collection("token")

	filter := bson.D{{"username", username}}
	_, err := collection.DeleteOne(context.Background(), filter)
	return err
}
