package pokeapi

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"ozum.safaoglu/pokemon-api/cache"
)

const (
	pokemonSpeciesCacheKeyPrefix = "pokemon-species-"
)

type cachedPokeAPI struct {
	pokeAPI PokeAPI
	cache   cache.Cache
}

// NewCachedPokeAPI returns a cached pokeapi instance
// CachedPokeAPI utilizes dependency injection to make use of a PokeAPI implementation
// and any cache implementation to search for a poke-api call's value in the cache repository first
// If it is found it will be returned, if it cannot be found PokeAPI implementation will be used to make a REST call
func NewCachedPokeAPI(pokeAPI PokeAPI, cache cache.Cache) PokeAPI {
	return &cachedPokeAPI{
		pokeAPI: pokeAPI,
		cache:   cache,
	}
}

func (c *cachedPokeAPI) generateCacheKey(name string) string {
	return pokemonSpeciesCacheKeyPrefix + name
}

func (c *cachedPokeAPI) marshalForCache(species *PokemonSpecies) (string, error) {
	bytes, err := json.Marshal(species)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (c *cachedPokeAPI) GetPokemons(ctx context.Context, offset int) (*Pokemons, error) {
	return c.pokeAPI.GetPokemons(ctx, offset)
}

func (c *cachedPokeAPI) GetPokemonSpecies(ctx context.Context, name string) (*PokemonSpecies, error) {
	val, err := c.cache.Get(ctx, c.generateCacheKey(name))
	if err != nil && err != redis.Nil {
		return nil, errors.Wrap(err, "Error while getting pokemon species cache")
	}
	if val != "" {
		var pokemonSpecies PokemonSpecies
		err := json.Unmarshal([]byte(val), &pokemonSpecies)
		if err != nil {
			logrus.Warn("Error occurred during unmarshalling pokemon species cache value", err)
		} else {
			return &pokemonSpecies, nil
		}
	}
	species, err := c.pokeAPI.GetPokemonSpecies(ctx, name)
	if err != nil {
		return nil, err
	}
	marshalled, err := c.marshalForCache(species)
	if err != nil {
		return nil, err
	}

	err = c.cache.Set(ctx, c.generateCacheKey(name), marshalled, time.Hour*48)
	if err != nil {
		logrus.Warnf("Error while setting the cache for pokemon species: %s", err)
	}

	return species, nil
}
