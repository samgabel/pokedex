package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Starts up a REPL session that takes in user input and returns the input.
func startRepl() {
	// create a new scanner
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		// `.Scan()` will block until `.Text()` is called
		// returns false when reached the end of input or an error
		if !scanner.Scan() {
			fmt.Println()
			break
		}

		// grab input and clean
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			fmt.Println()
			continue
		}
		inputCommand := input[0]

		// execute command
		command, exists := getCommands()[inputCommand]
		if exists {
			command.callback()
		} else {
			fmt.Printf("'%v' is an invalid command", inputCommand)
			fmt.Println()
		}

		// formatting
		fmt.Println()
	}
}

// This function is meant to be used to clean stdin text entering the REPL.
// See repl_test.go for tests against this function
func cleanInput(input string) []string {
	stringLower := strings.ToLower(input)
	return strings.Fields(stringLower)
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays the help message",
			callback:    helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exits the REPL program",
			callback:    exitCommand,
		},
	}
}

