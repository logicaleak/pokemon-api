package shakespeare

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"ozum.safaoglu/pokemon-api/cache"
)

const (
	translationPrefix = "translation-"
)

type cachedSPClient struct {
	client SPClient
	cache  cache.Cache
}

// NewCachedSPClient returns a cached pokeapi instance
// CachedPokeAPI utilizes dependency injection to make use of a PokeAPI implementation
// and any cache implementation to search for a poke-api call's value in the cache repository first
// If it is found it will be returned, if it cannot be found PokeAPI implementation will be used to make a REST call
func NewCachedSPClient(spClient SPClient, cache cache.Cache) SPClient {
	return &cachedSPClient{
		client: spClient,
		cache:  cache,
	}
}

func (c *cachedSPClient) generateHash(text string) (string, error) {
	h := md5.New()
	_, err := h.Write([]byte(text))
	if err != nil {
		return "", err
	}
	return string(h.Sum(nil)), nil
}

func (c *cachedSPClient) generateCacheKey(text string) (string, error) {
	textHash, err := c.generateHash(text)
	if err != nil {
		return "", err
	}
	return translationPrefix + string(textHash), nil
}

func (c *cachedSPClient) marshalForCache(translation *Translation) (string, error) {
	bytes, err := json.Marshal(translation)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (c *cachedSPClient) Translate(ctx context.Context, text string) (*Translation, error) {
	cacheKey, err := c.generateCacheKey(text)
	if err != nil {
		return nil, errors.Wrap(err, "Error while getting translation cache")
	}
	val, err := c.cache.Get(ctx, cacheKey)
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if val != "" {
		var translation Translation
		err := json.Unmarshal([]byte(val), &translation)
		if err != nil {
			logrus.Warn("Error occurred during unmarshalling pokemon species cache value", err)
		} else {
			return &translation, nil
		}
	}
	translation, err := c.client.Translate(ctx, text)
	if err != nil {
		return nil, err
	}

	marshaled, err := c.marshalForCache(translation)
	if err != nil {
		return nil, err
	}

	err = c.cache.Set(ctx, cacheKey, marshaled, time.Hour*48)
	if err != nil {
		logrus.Warnf("Error while setting the cache for translation: %s", err)
	}
	return translation, nil
}
