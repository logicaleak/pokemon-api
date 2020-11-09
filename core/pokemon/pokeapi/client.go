package pokeapi

import (
	"context"
	"encoding/json"
	"time"

	resty "github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// PokeAPI is the unimplemented pokeapi interface consisting of necessary endpoint methods
type PokeAPI interface {
	GetPokemonSpecies(context.Context, string) (*PokemonSpecies, error)
	GetPokemons(ctx context.Context, offset int) (*Pokemons, error)
}

type pokeAPIImpl struct {
	restyClient *resty.Client
	baseURI     string
}

// NewPokeAPI returns a new PokeAPI implementation
func NewPokeAPI(baseURI string) PokeAPI {
	return &pokeAPIImpl{
		restyClient: resty.New().SetRetryCount(3), // Default values of retry settings is good enough for our purposes
		baseURI:     baseURI,
	}
}

func (p *pokeAPIImpl) GetPokemons(ctx context.Context, offset int) (*Pokemons, error) {
	n := time.Now()
	logrus.Infof("Starting GET /pokemon")
	resp, err := p.restyClient.R().
		SetContext(ctx).
		EnableTrace().
		Get(p.baseURI + "/pokemon")
	if err != nil {
		return nil, errors.Wrap(err, "Error while requesting GET /pokemon")
	}
	var pokemons Pokemons
	err = json.Unmarshal(resp.Body(), &pokemons)
	if err != nil {
		return nil, err
	}

	duration := time.Since(n)
	logrus.WithField("duration", duration).Infof("Finished GET /pokemon without issues in %s", duration)

	return pokemons, nil
}

// GetPokemon returns the pokemon object from the poke-api
func (p *pokeAPIImpl) GetPokemonSpecies(ctx context.Context, name string) (*PokemonSpecies, error) {
	n := time.Now()
	logrus.Infof("Starting GET /pokemon-species/%s", name)

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

	duration := time.Since(n)
	logrus.WithField("duration", duration).Infof("Finished GET /pokemon-species/%s without issues in %s", name, duration)

	return &pokemonSpecies, nil
}
