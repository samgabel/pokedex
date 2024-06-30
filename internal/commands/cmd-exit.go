package commands

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config, param string) error {
	// exit command does not take additional inputs
	if param != "" {
		fmt.Println("Run 'exit' without command parameter")
		return nil
	}

	os.Exit(0)
	return nil
}
