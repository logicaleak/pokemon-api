package implementation

import (
	"context"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"ozum.safaoglu/pokemon-api/api/swagger/restapi/operations/pokemons"
	"ozum.safaoglu/pokemon-api/cache"
	"ozum.safaoglu/pokemon-api/config"
	"ozum.safaoglu/pokemon-api/core/pokemon/pokeapi"
	"ozum.safaoglu/pokemon-api/core/service"
	"ozum.safaoglu/pokemon-api/core/shakespeare"
)

type ShakespeareanPokemonAPI struct {
	spService service.ShakespeareanPokemon
}

func NewShakespeareanPokemonAPI() *ShakespeareanPokemonAPI {
	pokeAPI := pokeapi.NewPokeAPI(config.CFG.PokeapiBaseURL)
	spClient := shakespeare.NewClient(config.CFG.ShakespeareBaseURL)

	cache, err := cache.NewRedis(config.CFG.RedisAddr, "")
	if err != nil {
		panic(err)
	}
	cachedPokeAPI := pokeapi.NewCachedPokeAPI(pokeAPI, cache)
	cachedSPClient := shakespeare.NewCachedSPClient(spClient, cache)

	&ShakespeareanPokemonAPI{
		spService: service.NewShakespeareanPokemon(cachedPokeAPI, cachedSPClient, "en")
	} 
}

// GetPokemon implements the swagger api endpoint
// It returns the description of the passed pokemon in shakespearean English
func (s *ShakespeareanPokemonAPI) GetPokemon(params pokemons.GetParams) middleware.Responder {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	// defer cancel()
	// s.spService.GetDescription(ctx, params)

	return middleware.NotImplemented("operation pokemons.Get has not yet been implemented")
}
