package config

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret       string
	PageSize        int
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

var AppSetting = &App{}
var ServerSetting = &Server{}
var DatabaseSetting = &Database{}
var RedisSetting = &Redis{}
var RabbitMQSetting = &RabbitMQ{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
