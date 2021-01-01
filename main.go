package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"Food/config"
	"Food/helpers/database"
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

	gin.ForceConsoleColor()
	gin.SetMode(config.ServerSetting.RunMode)

	database, closeDB := database.NewDB(*config.DatabaseSetting)
	defer closeDB()
	database = config.SetupDB(database)

	// logging.NewLogger(*config.LoggerSetting)
	logging.NewZeroLog()

	config.SetupJWT(*config.AppSetting)
	config.SetupController(database)

	router := routers.InitRouter()

	readTimeout := config.ServerSetting.ReadTimeout
	writeTimeout := config.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%s", config.ServerSetting.HTTPPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	if config.ServerSetting.SSL {

		SSLKeys := &struct {
			CERT string
			KEY  string
		}{}

		//Generated using sh generate-certificate.sh
		SSLKeys.CERT = "./cert/myCA.cer"
		SSLKeys.KEY = "./cert/myCA.key"

		err := server.ListenAndServeTLS(SSLKeys.CERT, SSLKeys.KEY)
		if err != nil {
			log.Fatal("Web server (HTTPS): ", err)
		}
	} else {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Web server (HTTP): ", err)
		}
	}

}
