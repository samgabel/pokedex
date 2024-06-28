package commands

// Map values in the GetCommands return map
type cliCommand struct {
	name        string
	description string
	Callback    func(*Config) error
}

// This is a manifesto of all the commands available on the REPL.
// It shows the name, description, and actual callback function.
func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays the help message",
			Callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the REPL program",
			Callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Will display 20 location areas in the Pokemon world, with each subsequent call to `map` fetching the next 20",
			Callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Stands for 'map back', does what map does but in reverse",
			Callback:    commandMapb,
		},
	}
}
