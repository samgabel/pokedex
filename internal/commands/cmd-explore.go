package commands

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *Config, param string) error {
	// handle case where user doesn't eneter command param
	if param == "" {
		return errors.New("no command parameter entered for 'explore'")
	}

	// get encounterObj
	encounter, err := cfg.PokeapiClient.GetEncounters(param)
	if err != nil {
		fmt.Println("Location not found")
		return err
	}

	// print each pokemon name in the slice of structs
	fmt.Printf("Exploring %v...\n", param)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range encounter.PokemonEncounters {
		fmt.Println(" -", pokemon.Pokemon.Name)
	}

	return nil
}
