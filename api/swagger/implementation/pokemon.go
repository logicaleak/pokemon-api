package implementation

import (
	"github.com/go-openapi/runtime/middleware"
	"ozum.safaoglu/pokemon-api/api/swagger/restapi/operations/pokemons"
)

type ShakespeareanPokemonAPI struct {
}

func NewShakespeareanPokemonAPI() *ShakespeareanPokemonAPI {
	return nil
}

// GetPokemon implements the swagger api endpoint
// It returns the description of the passed pokemon in shakespearean English
func (s *ShakespeareanPokemonAPI) GetPokemon(params pokemons.GetParams) middleware.Responder {
	return middleware.NotImplemented("operation pokemons.Get has not yet been implemented")
}
