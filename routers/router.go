package routers

import (
  _ "Food/docs"
  "Food/helpers/jwt"
  "Food/middlewares"

  // "github.com/gin-gonic/contrib/static"
  "github.com/gin-gonic/gin"
  swaggerFiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initialize routing information
func InitRouter(jwtManager jwt.JwtManager) *gin.Engine {
  r := gin.New()
  r.Use(gin.Logger())
  r.Use(gin.Recovery())

  r.Use(middlewares.JWT(jwtManager))
  r.Use(middlewares.Security)

  // r.Use(static.Serve("/", static.LocalFile("./client/build", true)))

  r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

  InitRouterApi(r)

  return r
}
