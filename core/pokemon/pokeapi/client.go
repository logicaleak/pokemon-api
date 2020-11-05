package pokeapi

import (
	"encoding/json"

	resty "github.com/go-resty/resty/v2"
	"ozum.safaoglu/pokemon-api/config"
)

// PokeAPI is the unimplemented pokeapi interface consisting of necessary endpoint methods
type PokeAPI interface {
	GetPokemonSpecies(string) (*PokemonSpecies, error)
}

type pokeAPIImpl struct {
	restyClient *resty.Client
	baseURI     string
}

// NewPokeAPI returns a new PokeAPI implementation
func NewPokeAPI(config *config.Config) PokeAPI {
	return &pokeAPIImpl{
		restyClient: resty.New(),
		baseURI:     config.PokeAPIBaseURI.String(),
	}
}

// GetPokemon returns the pokemon object from the poke-api
func (p *pokeAPIImpl) GetPokemonSpecies(name string) (*PokemonSpecies, error) {
	resp, err := p.restyClient.R().
		EnableTrace().
		Get(p.baseURI + pokemon)
	if err != nil {
		return nil, err
	}

	var pokemonSpecies PokemonSpecies
	err = json.Unmarshal(resp.Body(), &pokemonSpecies)
	if err != nil {
		return nil, err
	}

	return &pokemonSpecies, nil
}
