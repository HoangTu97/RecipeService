package config

import (
	"Food/helpers/setting"
	"log"
	"time"

	"github.com/go-ini/ini"
)

var AppSetting = &setting.App{}
var ServerSetting = &setting.Server{}
var DatabaseSetting = &setting.Database{}
var RedisSetting = &setting.Redis{}
var RabbitMQSetting = &setting.RabbitMQ{}

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
