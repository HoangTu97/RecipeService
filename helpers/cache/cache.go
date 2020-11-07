package cache

type Cache interface {
	GenKey(data ...interface{}) string
	Set(key string, data interface{}, time int) error
	Exists(key string) bool
	Get(key string) ([]byte, error)
	Delete(key string) (bool, error)
	LikeDeletes(key string) error
}