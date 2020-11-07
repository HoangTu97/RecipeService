package cache

import (
	"Food/helpers/converter"
	"Food/helpers/setting"
	"encoding/json"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
)

type redisCache struct {
	redisConn *redis.Pool
}

func NewRedis(redisSetting setting.Redis) Cache {
	redisConn := &redis.Pool{
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

	return &redisCache{redisConn: redisConn}
}

func (r *redisCache) GenKey(data ...interface{}) string {
	values := make([]string, len(data))

	for i, dt := range data {
		values[i] = converter.ToStr(dt)
	}

	return strings.Join(values, "_")
}

// Set a key/value
func (r *redisCache) Set(key string, data interface{}, time int) error {
	conn := r.redisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

// Exists check a key
func (r *redisCache) Exists(key string) bool {
	conn := r.redisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get get a key
func (r *redisCache) Get(key string) ([]byte, error) {
	conn := r.redisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// Delete delete a kye
func (r *redisCache) Delete(key string) (bool, error) {
	conn := r.redisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes batch delete
func (r *redisCache) LikeDeletes(key string) error {
	conn := r.redisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = r.Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
