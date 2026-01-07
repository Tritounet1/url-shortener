package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"tidy/models"
	"tidy/utils"

	"github.com/gin-gonic/gin"
)

type createShortUrlDto struct {
	Url string
}

func createNewShortUrl(longUrl string) {
	coll := utils.Client.Database("db").Collection("url")

	random, err := utils.RandomStringCrypto(10)
	if err != nil {
		fmt.Println("Error generating random string:", err)
		return
	}
	fmt.Println("Random String:", random)

	url := models.NewUrl(longUrl, random)

	result, _ := coll.InsertOne(context.TODO(), url)

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
}

func createShortUrl(c *gin.Context) {
	body := createShortUrlDto{}
	data, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatusJSON(400, "Input format is wrong")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		c.AbortWithStatusJSON(400, "Can't match with the struct")
		return
	}

	createNewShortUrl(body.Url)

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
