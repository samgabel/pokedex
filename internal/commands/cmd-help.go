package commands

import (
	"fmt"
)

func commandHelp(cfg *Config) error {
	fmt.Println("-----------------------")
	fmt.Println("Welcome to Pokedex REPL")
	fmt.Println("-----------------------")
	for _, cmd := range GetCommands() {
		fmt.Println(cmd.name, "	", cmd.description)
	}

	return nil
}
