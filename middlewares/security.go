package middlewares

import (
	"Food/pkg/domain"
	"Food/dto/response"
	"Food/helpers/constants"

	"github.com/gin-gonic/gin"
)

var accessibleRoles map[string][]string

func init() {
	accessibleRoles = make(map[string][]string)
	// Security declare
  accessibleRoles["/api/private/user.*"] = []string{constants.ROLE.USER}
	accessibleRoles["/api/private/post"] = []string{constants.ROLE.USER}
	// Security declare end : dont remove
}

// Security is Security middleware
func Security(c *gin.Context) {
	roles, found := accessibleRoles[c.Request.URL.Path]
	if !found {
		c.Next()
		return
	}

	iUserInfo, exists := c.Get("UserInfo")
	if !exists {
		response.CreateErrorResponse(c, constants.ErrorStringApi.UNAUTHORIZED_ACCESS)
		c.Abort()
		return
	}

	userInfo := iUserInfo.(*domain.Token)
	if err := userInfo.Valid(); err != nil {
		response.CreateErrorResponse(c, constants.ErrorStringApi.UNAUTHORIZED_ACCESS)
		c.Abort()
		return
	}

	for _, authority := range roles {
		if !userInfo.HasAuthority(authority) {
			response.CreateErrorResponse(c, constants.ErrorStringApi.UNAUTHORIZED_ACCESS)
			c.Abort()
			return
		}
	}

	c.Next()
}

// Authenticated Authenticated
func Authenticated(c *gin.Context) {
	iUserInfo, exists := c.Get("UserInfo")
	if !exists {
		response.CreateErrorResponse(c, constants.ErrorStringApi.UNAUTHORIZED_ACCESS)
		c.Abort()
		return
	}

	userInfo := iUserInfo.(*domain.Token)
	if err := userInfo.Valid(); err != nil {
		response.CreateErrorResponse(c, constants.ErrorStringApi.UNAUTHORIZED_ACCESS)
		c.Abort()
		return
	}

	c.Next()
}

// HasAuthority HasAuthority
func HasAuthority(authority string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userInfo := c.MustGet("UserInfo").(*domain.Token)
		if !userInfo.HasAuthority(authority) {
			response.CreateErrorResponse(c, constants.ErrorStringApi.UNAUTHORIZED_ACCESS)
			c.Abort()
			return
		}
		c.Next()
	}
}
