package service

import (
	"context"
	"strings"

	"ozum.safaoglu/pokemon-api/core/pokemon/pokeapi"
	"ozum.safaoglu/pokemon-api/core/shakespeare"
)

// ShakespeareanPokemon is a service interface that returns a shakespearean description
// for the given pokemon name
type ShakespeareanPokemon interface {
	GetDescription(context.Context, string) (*PokemonDescription, error)
}

// NewShakespeareanPokemon returns the implementation of ShakespeareanPokemon interface
func NewShakespeareanPokemon(pokeAPI pokeapi.PokeAPI, spClient shakespeare.SPClient, lang string) ShakespeareanPokemon {
	return &defaultShakespeareanPokemonImpl{
		pokeAPI:  pokeAPI,
		spClient: spClient,
		lang:     lang,
	}
}

type defaultShakespeareanPokemonImpl struct {
	pokeAPI  pokeapi.PokeAPI
	spClient shakespeare.SPClient
	lang     string
}

func (d *defaultShakespeareanPokemonImpl) generateDescriptionFrom(species *pokeapi.PokemonSpecies) string {
	var b strings.Builder
	for _, e := range species.FlavourTextEntries {
		if e.Language.Name == d.lang && e.Version.Name == "ruby" {
			b.WriteString(e.FlavorText)
			b.WriteString(" ")
		}
	}
	return b.String()
}

// GetDescription first calls pokeapi to retrieve pokemon flavour descriptions and tries its best to combine something
// Afterwards it uses shakespeare translation api to obtain a translation for the decription set
func (d *defaultShakespeareanPokemonImpl) GetDescription(ctx context.Context, pokemon string) (*PokemonDescription, error) {
	species, err := d.pokeAPI.GetPokemonSpecies(ctx, pokemon)
	if err != nil {
		return nil, err
	}

	translation, err := d.spClient.Translate(ctx, d.generateDescriptionFrom(species))
	if err != nil {
		return nil, err
	}

	return &PokemonDescription{
		Name:        pokemon,
		Description: translation.Contents.Translated,
	}, nil
}
