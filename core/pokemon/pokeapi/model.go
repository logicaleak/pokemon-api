package pokeapi

type PokemonSpecies struct {
	FlavourTextEntries []FlavourTextEntries `json:"flavor_text_entries`
}

type FlavourTextEntries struct {
	FlavorText string `json:"flavor_text"`
}
