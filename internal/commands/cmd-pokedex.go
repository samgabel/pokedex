package commands

import "fmt"

func commandPokedex(cfg *Config, param string) error {
	// pokedex command does not take additional inputs
	if param != "" {
		fmt.Println("Run 'pokedex' without command parameter")
		return nil
	}

	// list pokedex entries
	fmt.Println("Your Pokedex:")
	pokedex := cfg.CaughtPokemon
	for entry := range pokedex {
		fmt.Println(" -", entry)
	}

	return nil
}
