package commands

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Config, param string) error {
	// handle case where user doesn't eneter command param
	if param == "" {
		return errors.New("no command parameter entered for 'catch'")
	}

	// get PokemonObj
	pokemon, err := cfg.PokeapiClient.GetPokemonStats(param)
	if err != nil {
		fmt.Println("Pokemon not found")
		return err
	}

	// attampt to catch pokemon
	fmt.Printf("Throwing a pokeball at %v...\n", pokemon.Name)

	// calculate catch number
	catchNum := rand.Intn(pokemon.BaseExperience) // lowest base exp of any pokemon is 36

	// determine catch
	if catchNum <= 40 {
		fmt.Printf("%v was CAUGHT! Adding %v to the pokedex.\n", pokemon.Name, pokemon.Name)
		cfg.CaughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}

	return nil
}
