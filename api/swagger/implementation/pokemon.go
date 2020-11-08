package implementation

import (
	"context"
	"time"

	"ozum.safaoglu/pokemon-api/api/swagger/models"
	"ozum.safaoglu/pokemon-api/api/swagger/restapi/operations/pokemondescription"
	"ozum.safaoglu/pokemon-api/cache"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"ozum.safaoglu/pokemon-api/config"
	"ozum.safaoglu/pokemon-api/core/pokemon/pokeapi"
	"ozum.safaoglu/pokemon-api/core/service"
	"ozum.safaoglu/pokemon-api/core/shakespeare"
)

type ShakespeareanPokemonAPI struct {
	spService service.ShakespeareanPokemon
}

func NewShakespeareanPokemonAPI() *ShakespeareanPokemonAPI {
	pokeAPI := pokeapi.NewPokeAPI(config.CFG.PokeapiBaseUrl)
	spClient := shakespeare.NewClient(config.CFG.ShakespeareBaseUrl)

	cache, err := cache.NewRedis(config.CFG.RedisAddr, "")
	if err != nil {
		panic(err)
	}
	cachedPokeAPI := pokeapi.NewCachedPokeAPI(pokeAPI, cache)
	cachedSPClient := shakespeare.NewCachedSPClient(spClient, cache)

	return &ShakespeareanPokemonAPI{
		spService: service.NewShakespeareanPokemon(cachedPokeAPI, cachedSPClient, "en"),
	}
}

// GetPokemon implements the swagger api endpoint
// It returns the description of the passed pokemon in shakespearean English
func (s *ShakespeareanPokemonAPI) GetPokemonDescription(params pokemondescription.GetPokemonNameParams) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	description, err := s.spService.GetDescription(ctx, params.Name)
	if err != nil {
		logrus.WithError(err).Errorf("Error while getting description of the pokemon")
		errStr := "An error occurred"
		return pokemondescription.NewGetPokemonNameDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errStr,
		})
	}

	return pokemondescription.NewGetPokemonNameOK().WithPayload(&models.Description{
		Description: description.Description,
		Name:        description.Name,
	})
}
