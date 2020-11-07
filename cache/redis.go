package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCacheImpl struct {
	client *redis.Client
}

// NewRedis returns a redis cache implementation
func NewRedis(url, password string) (Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       0,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	statusCmd := client.Ping(ctx)
	_, err := statusCmd.Result()
	if err != nil {
		return nil, err
	}
	return &redisCacheImpl{
		client: client,
	}, nil
}

func (r *redisCacheImpl) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *redisCacheImpl) Set(ctx context.Context, key string, val interface{}, ttl time.Duration) error {
	_, err := r.client.Set(ctx, key, val, ttl).Result()
	if err != nil {
		return err
	}
	return nil
}
