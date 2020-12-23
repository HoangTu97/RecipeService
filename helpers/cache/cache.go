package cache

import (
	"Food/helpers/converter"
	"Food/helpers/setting"
	"strings"
)

type Cache interface {
	GenKey(data ...interface{}) string
	Set(key string, data interface{}, time int) error
	Exists(key string) bool
	Get(key string) ([]byte, error)
	Delete(key string) (bool, error)
	LikeDeletes(key string) error
}

func genKey(data ...interface{}) string {
	values := make([]string, len(data))

	for i, dt := range data {
		values[i] = converter.ToStr(dt)
	}

	return strings.Join(values, "_")
}

func NewCache(cacheSetting setting.Cache) Cache {
	if cacheSetting.Type == "redis" {
		return NewRedis(cacheSetting)
	}
	if cacheSetting.Type == "memory" {
		return NewMem(cacheSetting)
	}
	return nil
}
