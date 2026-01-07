package routes

import (
	auth "tidy/routes/auth"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine) {

	router.GET("/:short_url", getShortUrl)

	router.POST("/auth/login", auth.Login)
	router.POST("/auth/register", auth.Register)

	// TODO: add Middleware for protect the routes bellow :
	router.POST("/url", createShortUrl)
}
