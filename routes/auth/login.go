package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"tidy/models"
	"tidy/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type loginDto struct {
	Username string
	Password string
}

func Login(c *gin.Context) {
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

	filter := bson.D{{Key: "username", Value: body.Username}}

	var user models.User

	coll := utils.Client.Database("db").Collection("user")

	err = coll.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.String(http.StatusNotFound, "Incorrect identifier.")
			return
		} else {
			panic(err)
		}
	}

	if utils.VerifyPassword(body.Password, user.Password) {
		token, err := utils.CreateJWTToken(body.Username)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"message": "Error while creating JWT token.",
				"success": false,
			})
		} else {
			utils.SaveToken(user.Username, token)
			c.JSON(http.StatusOK, gin.H{
				"message": "User connected.",
				"token":   token,
				"success": true,
			})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Incorrect identifier.",
			"success": false,
		})
	}
}
