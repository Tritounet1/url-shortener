package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"tidy/models"
	"tidy/utils"

	"github.com/gin-gonic/gin"
)

type registerDto struct {
	Username string
	Password string
}

func createUser(username string, password string) {
	coll := utils.Client.Database("db").Collection("user")

	hashedPassword, _ := utils.HashPassword(password)

	user := models.NewUser(username, hashedPassword)

	result, _ := coll.InsertOne(context.TODO(), user)

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
}

func Register(c *gin.Context) {
	body := registerDto{}

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

	createUser(body.Username, body.Password)

	c.JSON(http.StatusOK, gin.H{
		"message": "New user create",
		"success": true,
	})
}
