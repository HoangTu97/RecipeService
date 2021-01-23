package routers

import (
	"Food/config"

	"github.com/gin-gonic/gin"
)

func registerPrivateApi(apiRoutes *gin.RouterGroup) {
	privateRoutes := apiRoutes.Group("/private")
	// Api declare
	{
		privatePostRoutes := privateRoutes.Group("/post")
		privatePostRoutes.POST("", config.PostController.CreatePost)
	}
	// Api declare end : dont remove
}