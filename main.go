package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"tidy/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	client *mongo.Client
)

func initClient() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")

	client, err = mongo.Connect(options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}
}

func randomStringCrypto(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
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
func createDatabase(databaseName string, collections []string) {
	for _, collection := range collections {
		client.Database(databaseName).CreateCollection(context.Background(), collection)
	}
}

func createNewShortUrl(longUrl string) {
	// Select a collection
	coll := client.Database("db").Collection("url")

	random, err := randomStringCrypto(10)
	if err != nil {
		fmt.Println("Error generating random string:", err)
		return
	}
	fmt.Println("Random String:", random)

	url := models.NewUrl(longUrl, random)

	// Insert the new document
	result, _ := coll.InsertOne(context.TODO(), url)

	// Print the ID of the document (automatically create by mongo)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
}

func router() {

}

func main() {
	initClient()

	collections := []string{"user", "url"}

	createDatabase("db", collections)

	// Close client connection if the program crash or is terminate by force
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// createNewShortUrl("https://google.com")

	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	router.POST("/auth/login", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/auth/register", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/url", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/:short_url", func(c *gin.Context) {
		shortUrl := c.Param("short_url")

		filter := bson.D{{Key: "short_url", Value: shortUrl}}

		var result models.Url

		coll := client.Database("db").Collection("url")

		err := coll.FindOne(context.TODO(), filter).Decode(&result)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				fmt.Println("No documents found")
				c.String(http.StatusNotFound, "URL not found")
				return
			} else {
				panic(err)
			}
		}

		// Update the click counter
		updateOpts := options.UpdateOne().SetUpsert(false)
		update := bson.D{{"$inc", bson.D{{"total_clicks", 1}}}}

		_, err = coll.UpdateOne(context.TODO(), filter, update, updateOpts)

		if err != nil {
			panic(err)
		}

		c.Redirect(http.StatusMovedPermanently, result.LongUrl)
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	router.Run()

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
