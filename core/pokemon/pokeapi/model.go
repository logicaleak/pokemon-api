package pokeapi

// PokemonSpecies is the returned model from the species endpoint
type PokemonSpecies struct {
	FlavourTextEntries []FlavourTextEntries `json:"flavor_text_entries"`
}

// FlavourTextEntries is the description entries for the pokemon
type FlavourTextEntries struct {
	FlavorText string   `json:"flavor_text"`
	Language   Language `json:"language"`
	Version    Version  `json:"version"`
}

// Language for flavour text entry
type Language struct {
	Name string `json:"name"`
}

type Version struct {
	Name string `json:"name"`
}
