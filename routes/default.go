package routes

import (
	"context"
	"fmt"
	"net/http"
	"tidy/models"
	"tidy/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func getShortUrl(c *gin.Context) {
	shortUrl := c.Param("short_url")

	filter := bson.D{{Key: "short_url", Value: shortUrl}}

	var result models.Url

	coll := utils.Client.Database("db").Collection("url")

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
}
