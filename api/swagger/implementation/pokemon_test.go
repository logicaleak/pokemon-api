package implementation

import (
	"errors"
	"reflect"
	"testing"

	"github.com/go-openapi/runtime/middleware"
	"github.com/stretchr/testify/mock"
	"ozum.safaoglu/pokemon-api/api/swagger/models"
	"ozum.safaoglu/pokemon-api/api/swagger/restapi/operations/pokemondescription"
	"ozum.safaoglu/pokemon-api/core/service"
)

func TestShakespeareanPokemonAPI_GetPokemonDescription(t *testing.T) {
	pokemon := "pikachu"
	pokemon2 := "charizard"
	description := "pikachu-desc"
	errStr := "An error occurred"
	mockService := &service.MockShakespeareanPokemon{}
	mockService.On("GetDescription", mock.Anything, pokemon).Return(&service.PokemonDescription{
		Description: description,
		Name:        pokemon,
	}, nil)
	mockService.On("GetDescription", mock.Anything, pokemon2).Return(nil, errors.New("Some error"))
	type args struct {
		params pokemondescription.GetPokemonNameParams
	}
	tests := []struct {
		name string
		s    *ShakespeareanPokemonAPI
		args args
		want middleware.Responder
	}{
		{
			name: "standard",
			s: &ShakespeareanPokemonAPI{
				spService: mockService,
			},
			args: args{
				params: pokemondescription.GetPokemonNameParams{
					Name: pokemon,
				},
			},
			want: pokemondescription.NewGetPokemonNameOK().WithPayload(&models.Description{
				Description: description,
				Name:        pokemon,
			}),
		},
		{
			name: "error",
			s: &ShakespeareanPokemonAPI{
				spService: mockService,
			},
			args: args{
				params: pokemondescription.GetPokemonNameParams{
					Name: pokemon2,
				},
			},
			want: pokemondescription.NewGetPokemonNameDefault(500).WithPayload(&models.Error{
				Code:    500,
				Message: &errStr,
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetPokemonDescription(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShakespeareanPokemonAPI.GetPokemonDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}
