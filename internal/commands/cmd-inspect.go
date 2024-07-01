package commands

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *Config, param string) error {
	// handle case where user doesn't eneter command param
	if param == "" {
		return errors.New("no command parameter entered for 'inspect'")
	}

	// process requested pokemon
	entry, ok := cfg.CaughtPokemon[param]

	// deal with either not found or not caught yet
	if !ok {
		pokemon, err := cfg.PokeapiClient.GetPokemonStats(param)
		if err != nil {
			fmt.Println("Pokemon not found")
			return err
		}

		fmt.Printf("You have not caught %v yet", pokemon.Name)
		fmt.Println()
		return nil
	}

	// print out pokedex entry of caught requested pokemon
	fmt.Printf("Name: %v\n", entry.Name)
	fmt.Printf("Height: %v\n", entry.Height)
	fmt.Printf("Weight: %v\n", entry.Weight)
	fmt.Println("Stats:")
	for _, stat := range entry.Stats {
		fmt.Printf(" - %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range entry.Types {
		fmt.Printf(" - %v\n", typ.Type.Name)
	}

	return nil
}
