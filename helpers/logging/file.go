package logging

import (
	"Food/config"
	"fmt"
	"time"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", config.AppSetting.RuntimeRootPath, config.AppSetting.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		config.AppSetting.LogSaveName,
		time.Now().Format(config.AppSetting.TimeFormat),
		config.AppSetting.LogFileExt,
	)
}
