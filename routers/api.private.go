package routers

import (
	"Food/config"
	"Food/helpers/constants"
	"Food/middlewares"

	"github.com/gin-gonic/gin"
)

func registerPrivateApi(apiRoutes *gin.RouterGroup) {
	privateRoutes := apiRoutes.Group("/private")
	privateRoutes.Use(middlewares.Authenticated)
	{
		privatePostRoutes := privateRoutes.Group("/post")
		privatePostRoutes.Use(middlewares.HasAuthority(constants.ROLE_USER))
		privatePostRoutes.POST("", config.PostController.CreatePost)
	}
}