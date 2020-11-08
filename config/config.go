package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config is the app configuration object
type Config struct {
	PokeapiBaseUrl     string
	ShakespeareBaseUrl string
	RedisAddr          string
}

// CFG Global config of the app
var CFG Config

func init() {
	env := os.Getenv("ENV")
	fmt.Println(env)
	godotenv.Load(".env." + env)

	CFG = Config{
		PokeapiBaseUrl:     os.Getenv("POKEMONAPI_POKEAPI_BASE_URL"),
		ShakespeareBaseUrl: os.Getenv("POKEMONAPI_POKEAPI_SHAKESPEARE_BASE_URL"),
		RedisAddr:          os.Getenv("POKEMONAPI_REDIS_ADDR"),
	}

	fmt.Println(CFG)
}
