package implementation

import (
	"context"
	"time"

	"ozum.safaoglu/pokemon-api/api/swagger/models"
	"ozum.safaoglu/pokemon-api/api/swagger/restapi/operations/pokemondescription"
	"ozum.safaoglu/pokemon-api/api/swagger/restapi/operations/pokemons"
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
func (s *ShakespeareanPokemonAPI) GetPokemonDescription(params pokemondescription.GetV1PokemonPokemonNameParams) middleware.Responder {
	start := time.Now()
	logrus.Infof("Executing GetPokemonDescription for pokemon %s", params.PokemonName)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	description, err := s.spService.GetDescription(ctx, params.PokemonName)
	if err != nil {
		logrus.WithError(err).Errorf("Error while getting description of the pokemon")
		errStr := "An error occurred"
		return pokemondescription.NewGetPokemonNameDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errStr,
		})
	}

	elapsed := time.Since(start)
	logrus.Infof("Execution of GetPokemonDescription finished in %s", elapsed)
	return pokemondescription.NewGetPokemonNameOK().WithPayload(&models.Description{
		Description: description.Description,
		Name:        description.Name,
	})
}

func (s *ShakespeareanPokemonAPI) GetPokemons(params pokemons.GetV1PokemonParams) middleware.Responder {
	start := time.Now()
	logrus.Infof("Executing GetPokemons")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	offset := 0
	if params.Offset != nil {
		offset = int(*params.Offset)
	}
	pokemonsResp, err := s.spService.GetPokemons(ctx, offset)
	if err != nil {
		logrus.WithError(err).Errorf("Error while getting pokemons")
		errStr := "An error occurred"
		return pokemons.NewGetV1PokemonDefault(500).WithPayload(&models.Error{
			Code:    500,
			Message: &errStr,
		})
	}

	elapsed := time.Since(start)
	logrus.Infof("Execution of GetPokemons finished in %s", elapsed)

	var pokemonsAPIResp models.Pokemons
	pokemonsAPIResp.Count = int64(pokemonsResp.Count)
	var results []*models.Result
	for _, r := range pokemonsResp.Results {
		results = append(results, &models.Result{
			Name: r.Name,
		})
	}
	pokemonsAPIResp.Results = results
	return pokemons.NewGetV1PokemonOK().WithPayload(&pokemonsAPIResp)
}
