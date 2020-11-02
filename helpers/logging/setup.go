package logging

import (
	"Food/helpers/file"
	"Food/helpers/setting"
	"log"
)

// Setup initialize the log instance
func Setup(appSetting setting.App) {
	var err error
	filePath := getLogFilePath(appSetting.RuntimeRootPath, appSetting.LogSavePath)
	fileName := getLogFileName(appSetting.LogSaveName, appSetting.TimeFormat, appSetting.LogFileExt)
	F, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}