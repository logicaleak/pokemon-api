package service

import (
	"context"
	"reflect"
	"testing"

	"ozum.safaoglu/pokemon-api/core/pokemon/pokeapi"
	"ozum.safaoglu/pokemon-api/core/shakespeare"
)

func Test_defaultShakespeareanPokemonImpl_generateDescriptionFrom(t *testing.T) {
	type fields struct {
		pokeAPI  pokeapi.PokeAPI
		spClient shakespeare.SPClient
		lang     string
	}
	type args struct {
		species *pokeapi.PokemonSpecies
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "standard",
			fields: fields{
				pokeAPI:  nil,
				spClient: nil,
				lang:     "en",
			},
			args: args{
				species: &pokeapi.PokemonSpecies{
					FlavourTextEntries: []pokeapi.FlavourTextEntries{
						{
							Language: pokeapi.Language{
								Name: "en",
							},
							FlavorText: "First sentence of the entry.",
							Version: pokeapi.Version{
								Name: "ruby",
							},
						},
						{
							Language: pokeapi.Language{
								Name: "en",
							},
							FlavorText: "Second sentence of the entry.",
							Version: pokeapi.Version{
								Name: "ruby",
							},
						},
					},
				},
			},
			want: "First sentence of the entry. Second sentence of the entry. ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &defaultShakespeareanPokemonImpl{
				pokeAPI:  tt.fields.pokeAPI,
				spClient: tt.fields.spClient,
				lang:     tt.fields.lang,
			}
			if got := d.generateDescriptionFrom(tt.args.species); got != tt.want {
				t.Errorf("defaultShakespeareanPokemonImpl.generateDescriptionFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_defaultShakespeareanPokemonImpl_GetDescription(t *testing.T) {
	ctx := context.Background()
	flavorTextOne := "First sentence of description."
	flavorTextTwo := "Second sentence of description."
	translation := "Translation first."
	mPokeapi := &pokeapi.MockPokeAPI{}
	mPokeapi.On("GetPokemonSpecies", ctx, "pikachu").Return(&pokeapi.PokemonSpecies{
		FlavourTextEntries: []pokeapi.FlavourTextEntries{
			{
				Language: pokeapi.Language{
					Name: "en",
				},
				FlavorText: flavorTextOne,
				Version: pokeapi.Version{
					Name: "ruby",
				},
			},
			{
				Language: pokeapi.Language{
					Name: "en",
				},
				FlavorText: flavorTextTwo,
				Version: pokeapi.Version{
					Name: "ruby",
				},
			},
		},
	}, nil)
	mspClient := &shakespeare.MockSPClient{}
	mspClient.On("Translate", ctx, flavorTextOne+" "+flavorTextTwo+" ").Return(&shakespeare.Translation{
		Success: shakespeare.Success{Total: 1},
		Contents: shakespeare.Content{
			Text:       flavorTextOne + " " + flavorTextTwo + "  ",
			Translated: translation,
		},
	}, nil)

	type fields struct {
		pokeAPI  pokeapi.PokeAPI
		spClient shakespeare.SPClient
		lang     string
	}
	type args struct {
		ctx     context.Context
		pokemon string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PokemonDescription
		wantErr bool
	}{
		{
			name: "standard",
			fields: fields{
				pokeAPI:  mPokeapi,
				spClient: mspClient,
				lang:     "en",
			},
			args: args{
				ctx:     ctx,
				pokemon: "pikachu",
			},
			want: &PokemonDescription{
				Name:        "pikachu",
				Description: translation,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &defaultShakespeareanPokemonImpl{
				pokeAPI:  tt.fields.pokeAPI,
				spClient: tt.fields.spClient,
				lang:     tt.fields.lang,
			}
			got, err := d.GetDescription(tt.args.ctx, tt.args.pokemon)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultShakespeareanPokemonImpl.GetDescription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultShakespeareanPokemonImpl.GetDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}
