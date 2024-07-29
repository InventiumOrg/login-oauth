package middlewares

import (
	"fmt"
	"login-oauth/config"
	"login-oauth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateAuthentication(c *gin.Context) {

	jwt, err := c.Cookie("token")

	if err != nil {
		c.Status(http.StatusUnauthorized)
		c.Abort()
	}
	config, err := config.LoadConfig(".")
	actualToken, err := utils.ValidateToken(jwt, config.PublicKey)

	if err != nil {
		c.Status(http.StatusUnauthorized)
		c.Abort()
	}
	fmt.Println(actualToken)
	c.Next()
}
