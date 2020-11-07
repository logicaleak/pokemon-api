package cache

import (
	"context"
	"time"
)

// Cache is expected to save and retrieve string type values to a cache repository
type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, val interface{}, ttl time.Duration) error
}
