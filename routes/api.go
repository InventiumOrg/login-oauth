package routes

import (
	"fmt"
	"login-oauth/handlers"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.New()
	// router.Use(middlewares.ValidateAuthentication)
	router.GET("/signout", func(c *gin.Context) {
		fmt.Println("Login Success")
	})

	router.POST("/signin", func(c *gin.Context) {
		fmt.Println("Login Success")
	})

	router.POST("/signup", func(c *gin.Context) {
		handlers.SignUp(c)
	})

	return router
}
