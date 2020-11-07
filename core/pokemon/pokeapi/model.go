package pokeapi

// PokemonSpecies is the returned model from the species endpoint
type PokemonSpecies struct {
	FlavourTextEntries []FlavourTextEntries `json:"flavor_text_entries"`
}

// FlavourTextEntries is the description entries for the pokemon
type FlavourTextEntries struct {
	FlavorText string `json:"flavor_text"`
}
