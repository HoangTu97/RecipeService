package middlewares

import (
	"Food/helpers/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if len(token) == 0 {
		c.Next()
		return
	}
	token = strings.Fields(token)[1]

	parsedToken, err := jwt.ParseToken(token)
	if err != nil {
		c.Next()
		return
	}

	c.Set("UserInfo", parsedToken)

	c.Next()
}
