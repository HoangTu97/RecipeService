package gredis

import (
	"Food/helpers/setting"
	"time"

	"github.com/gomodule/redigo/redis"
)

// Setup Initialize the Redis instance
func Setup(redisSetting setting.Redis) error {
	RedisConn = &redis.Pool{
		MaxIdle:     redisSetting.MaxIdle,
		MaxActive:   redisSetting.MaxActive,
		IdleTimeout: redisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisSetting.Host)
			if err != nil {
				return nil, err
			}
			if redisSetting.Password != "" {
				if _, err := c.Do("AUTH", redisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}