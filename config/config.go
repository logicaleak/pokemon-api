package config

import "net/url"

type Config struct {
	PokeAPIBaseURI url.URL
	ShakespeareURI url.URL
}
