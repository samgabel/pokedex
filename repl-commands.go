package main

import (
	"fmt"
	"os"
)

func helpCommand() error {
	fmt.Println("-----------------------")
	fmt.Println("Welcome to Pokedex REPL")
	fmt.Println("-----------------------")
	for _, cmd := range getCommands() {
		fmt.Println(cmd.name, "	", cmd.description)
	}

	return nil
}

func exitCommand() error {
	os.Exit(0)
	return nil
}

