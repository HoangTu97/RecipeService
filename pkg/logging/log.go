package logging

import (
	"Food/helpers/file"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

type Level int

var (
	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
}

type logger struct {
	logger *log.Logger
}

func NewLogger(config Config) Logger {
	filePath := getLogFilePath(config.RuntimeRootPath, config.LogSavePath)
	fileName := getLogFileName(config.LogSaveName, config.TimeFormat, config.LogFileExt)
	file, err := file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	log := log.New(file, DefaultPrefix, log.LstdFlags)

	return &logger{ logger: log }
}

// Debug output logs at debug level
func (log *logger) Debug(v ...interface{}) {
	log.setPrefix(DEBUG)
	log.logger.Println(v...)
	fmt.Println(v...)
}

// Info output logs at info level
func (log *logger) Info(v ...interface{}) {
	log.setPrefix(INFO)
	log.logger.Println(v...)
	fmt.Println(v...)
}

// Warn output logs at warn level
func (log *logger) Warn(v ...interface{}) {
	log.setPrefix(WARNING)
	log.logger.Println(v...)
	fmt.Println(v...)
}

// Error output logs at error level
func (log *logger) Error(v ...interface{}) {
	log.setPrefix(ERROR)
	log.logger.Println(v...)
	fmt.Println(v...)
}

// Fatal output logs at fatal level
func (log *logger) Fatal(v ...interface{}) {
	log.setPrefix(FATAL)
	log.logger.Fatalln(v...)
}

// setPrefix set the prefix of the log output
func (log *logger) setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	log.logger.SetPrefix(logPrefix)
}