package setting

import "time"

type App struct {
	JwtSecret       string
	PageSize        int
}

type Logger struct {
	RuntimeRootPath string
	LogSavePath     string
	LogSaveName     string
	LogFileExt      string
	TimeFormat      string
}

type Server struct {
	RunMode      string
	HTTPPort     string
	SSL          bool
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	Host     string
	Port     string
	Type     string
	Name     string
	User     string
	Password string
}

type Redis struct {
	Host        string
	Port        string
	Password    string
	SSL         bool
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type RabbitMQ struct {
	Host     string
	Port     string
	User     string
	Password string
}