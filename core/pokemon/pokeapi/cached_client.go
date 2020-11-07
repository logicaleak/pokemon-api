package pokeapi

import (
	"context"
	"encoding/json"

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

func (c *cachedPokeAPI) GetPokemonSpecies(ctx context.Context, name string) (*PokemonSpecies, error) {
	val, err := c.cache.Get(ctx, pokemonSpeciesCacheKeyPrefix+name)
	if err != nil {
		return nil, err
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
	return c.pokeAPI.GetPokemonSpecies(ctx, name)
}
