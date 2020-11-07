package pokeapi

import (
	"context"
	"encoding/json"

	resty "github.com/go-resty/resty/v2"
)

// PokeAPI is the unimplemented pokeapi interface consisting of necessary endpoint methods
type PokeAPI interface {
	GetPokemonSpecies(context.Context, string) (*PokemonSpecies, error)
}

type pokeAPIImpl struct {
	restyClient *resty.Client
	baseURI     string
}

// NewPokeAPI returns a new PokeAPI implementation
func NewPokeAPI(baseURI string) PokeAPI {
	return &pokeAPIImpl{
		restyClient: resty.New(),
		baseURI:     baseURI,
	}
}

// GetPokemon returns the pokemon object from the poke-api
func (p *pokeAPIImpl) GetPokemonSpecies(ctx context.Context, name string) (*PokemonSpecies, error) {
	resp, err := p.restyClient.R().
		SetContext(ctx).
		EnableTrace().
		Get(p.baseURI + pokemonSpecies + "/" + name)
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
