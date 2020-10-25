package routers

import (
	_ "Food/docs"
	"Food/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// r.Use(include.CORS())

	r.Use(middlewares.JWT)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	InitRouterApi(r)

	return r
}
