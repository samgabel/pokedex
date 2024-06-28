package main

import (
	"github.com/samgabel/pokedex/internal/commands"
)

func main() {
	cfg := &commands.Config{}
	startRepl(cfg)
}
