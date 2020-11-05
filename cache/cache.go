package cache

type Cache interface {
	Get(key string) ([]byte, error)
}
