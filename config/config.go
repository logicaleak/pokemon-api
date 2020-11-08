package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config is the app configuration object
type Config struct {
	PokeapiBaseURL     string
	ShakespeareBaseURL string
	RedisAddr          string
}

// CFG Global config of the app
var CFG *Config

func init() {
	err := envconfig.Process("pokemonapi", CFG)
	if err != nil {
		panic(err)
	}
}
