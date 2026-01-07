package main

import (
	"context"
	"tidy/routes"
	"tidy/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitClient()

	collections := []string{"user", "url", "visitor"}

	utils.CreateDatabase("db", collections)

	// Close client connection if the program crash or is terminate with force
	defer func() {
		if err := utils.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	router := gin.Default()

	routes.AddRoutes(router)

	router.Run()
}
