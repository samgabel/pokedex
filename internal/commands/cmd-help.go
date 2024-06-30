package commands

import (
	"fmt"
)

func commandHelp(cfg *Config, param string) error {
	// help command does not take additional inputs
	if param != "" {
		fmt.Println("Run 'help' without command parameter")
		return nil
	}

	fmt.Println("-----------------------")
	fmt.Println("Welcome to Pokedex REPL")
	fmt.Println("-----------------------")
	for _, cmd := range GetCommands() {
		fmt.Println(cmd.name, cmd.description)
	}

	return nil
}
