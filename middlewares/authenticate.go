package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateAuthentication(c *gin.Context) {
	jwt, err := c.Cookie("session")
	if err != nil {
		c.Status(http.StatusUnauthorized)
		c.Abort()
	}
	fmt.Println(jwt)
	c.Next()
}
