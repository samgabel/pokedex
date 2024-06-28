package commands

import "github.com/samgabel/pokedex/internal/pokeapi"

// Commands will accept an argument of a pointer to this struct
type Config struct {
	PokeapiClient       pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
}
