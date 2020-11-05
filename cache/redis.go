package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCacheImpl struct {
	client *redis.Client
}

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
