package routes

import (
	"fmt"
	"login-oauth/handlers"
	"login-oauth/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.New()

	router.GET("/signout", func(c *gin.Context) {
		fmt.Println("Login Success")
	})

	router.POST("/signin", func(c *gin.Context) {
		handlers.Signin(c)
	})

	router.POST("/signup", func(c *gin.Context) {
		handlers.SignUp(c)
	})

	router.Use(middlewares.ValidateAuthentication)

	router.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "User is authenticated"})
	})
	return router
}
