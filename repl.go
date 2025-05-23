package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/samgabel/pokedex/internal/commands"
	"github.com/samgabel/pokedex/internal/pokeapi"
)

const (
	red    string = "\033[31m"       // ANSI escape code for red
	yellow string = "\033[38;5;226m" // ANSI escape code for yellow
	reset  string = "\033[0m"        // Reset color
)

// Starts up a REPL session
func startRepl(cfg *commands.Config) {

	// create a new scanner
	scanner := bufio.NewScanner(os.Stdin)

	// create new pokeapi.Client (http.Client and pokecache.Cache)
	cfg.PokeapiClient = pokeapi.NewClient(time.Second*5, time.Minute*3)

	// create a map to store "caught" pokemon for the pokedex
	cfg.CaughtPokemon = make(map[string]pokeapi.PokemonObj)

	// enter into REPL (infinite loop)
	for {
		fmt.Printf("%sPokedex > %s", red, reset)

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

		// handle additional input for certain commands
		commandParameter := ""
		if len(input) > 1 {
			commandParameter = input[1]
		}

		// handle too many inputs
		if len(input) > 2 {
			fmt.Println("Too many inputs try again")
			fmt.Println()
			continue
		}

		// execute command
		command, exists := commands.GetCommands()[inputCommand]
		if exists {
			command.Callback(cfg, commandParameter)
		} else {
			fmt.Printf("%s%v%s is an invalid command", yellow, inputCommand, reset)
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
