package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"Food/config"
	"Food/helpers/cache"
	"Food/helpers/database"
	"Food/helpers/jwt"
	"Food/helpers/logging"
	"Food/routers"
)

// @title Food API
// @version 1.0
// @description An example of gin
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config.Setup()
	database, closeDB := database.NewDB(*config.DatabaseSetting)
	defer closeDB()
	database = config.SetupDB(database)
	logging.NewLogger(*config.LoggerSetting)
	cache.NewRedis(*config.RedisSetting)
	jwt.Setup(*config.AppSetting)
	config.SetupController(database)

	gin.SetMode(config.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := config.ServerSetting.ReadTimeout
	writeTimeout := config.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%s", config.ServerSetting.HTTPPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

	// if config.ServerSetting.SSL {

	// 	SSLKeys := &struct {
	// 		CERT string
	// 		KEY  string
	// 	}{}

	// 	//Generated using sh generate-certificate.sh
	// 	SSLKeys.CERT = "./cert/myCA.cer"
	// 	SSLKeys.KEY = "./cert/myCA.key"

	// 	routersInit.RunTLS(":"+config.ServerSetting.HTTPPort, SSLKeys.CERT, SSLKeys.KEY)
	// } else {
	// 	routersInit.Run(":" + config.ServerSetting.HTTPPort)
	// }

}
