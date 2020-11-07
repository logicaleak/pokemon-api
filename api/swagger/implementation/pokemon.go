package implementation

import (
	"github.com/go-openapi/runtime/middleware"
	"ozum.safaoglu/pokemon-api/api/swagger/restapi/operations/pokemons"
)

// GetPokemon implements the swagger api endpoint
// It returns the description of the passed pokemon in shakespearean English
func GetPokemon(params pokemons.GetParams) middleware.Responder {
	return middleware.NotImplemented("operation pokemons.Get has not yet been implemented")
}
