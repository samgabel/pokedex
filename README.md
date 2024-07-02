# Pokedex (go)




## Overview

This is a REPL (Read-Eval-Print Loop) built in Go. It takes user input and interacts with the [PokiAPI](https://pokeapi.co/) in order to simulate "command-line exploration" of the pokemon world. I used this project to better familiarize myself with Go's concurrency patterns, as well as data serialization, HTTP networking, and caching.

![showcase](https://github.com/samgabel/pokedex/blob/main/showcase.png?raw=true)

My hope is that I can continue to come back to this project to work on new ideas as my software design skills improve. Below is a list of current ideas I have for the project.




## Usage

1. Clone this repository to your client machine
2. Run `go run .` in the project root or compile and run the binary to enter into the REPL
3. Once in the REPL type in 'help' for a list of commands




## TODO

- [ ] Update the CLI to support the "up" arrow to cycle through previous commands
- [ ] Simulate battles between pokemon
- [ ] Keep pokemon in a "party" and allow them to level up
- [ ] Allow for pokemon that are caught to evolve after a set amount of time
- [ ] Persist a user's Pokedex to disk so they can save progress between sessions
- [ ] Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
- [ ] Random encounters with wild pokemon
- [ ] Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon
