package pokeapi

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"ozum.safaoglu/pokemon-api/cache"
)

func Test_cachedPokeAPI_GetPokemonSpecies_ChecksCacheFirst(t *testing.T) {
	pokemon := "pikachu"
	pokemonSpecies := &PokemonSpecies{
		FlavourTextEntries: []FlavourTextEntries{
			{
				FlavorText: "Pikachu is an electric pokemon",
			},
		},
	}
	marshalledPokemonSpecies, err := json.Marshal(pokemonSpecies)
	assert.Nil(t, err)
	ctx := context.Background()

	cache := &cache.MockCache{}
	cache.On("Get", ctx, pokemonSpeciesCacheKeyPrefix+pokemon).Return(string(marshalledPokemonSpecies), nil)

	pokeAPI := &MockPokeAPI{}

	cachedPokeAPI := NewCachedPokeAPI(pokeAPI, cache)
	species, err := cachedPokeAPI.GetPokemonSpecies(ctx, pokemon)
	assert.Nil(t, err)
	pokeAPI.AssertNotCalled(t, "GetPokemonSpecies")
	assert.Equal(t, pokemonSpecies, species)
}

func Test_cachedPokeAPI_GetPokemonSpecies_CallsAPI_WithoutCache(t *testing.T) {
	pokemon := "pikachu"
	pokemonSpecies := &PokemonSpecies{
		FlavourTextEntries: []FlavourTextEntries{
			{
				FlavorText: "Pikachu is an electric pokemon",
			},
		},
	}

	ctx := context.Background()

	cache := &cache.MockCache{}
	cache.On("Get", ctx, pokemonSpeciesCacheKeyPrefix+pokemon).Return("", nil)
	cache.On("Set", ctx, pokemonSpeciesCacheKeyPrefix+pokemon, mock.Anything, mock.Anything).Return(nil).Once()

	pokeAPI := &MockPokeAPI{}
	pokeAPI.On("GetPokemonSpecies", ctx, pokemon).Return(pokemonSpecies, nil)

	cachedPokeAPI := NewCachedPokeAPI(pokeAPI, cache)
	species, err := cachedPokeAPI.GetPokemonSpecies(ctx, pokemon)
	assert.Nil(t, err)
	pokeAPI.AssertNumberOfCalls(t, "GetPokemonSpecies", 1)
	assert.Equal(t, pokemonSpecies, species)

	cache.AssertExpectations(t)
}
