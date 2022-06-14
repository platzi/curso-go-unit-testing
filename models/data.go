package models

type PokeApiPokemonResponse struct {
	Id          int           `json:"id"`
	Name        string        `json:"name"`
	PokemonType []PokemonType `json:"types"`
	Stats       []PokemonStat `json:"stats"`
}

type PokemonType struct {
	Slot    int      `json:"slot"`
	RefType BaseName `json:"type"`
}

type PokemonStat struct {
	BaseStat int      `json:"base_stat"`
	Stat     BaseName `json:"stat"`
}

type BaseName struct {
	Name string `json:"name"`
}
