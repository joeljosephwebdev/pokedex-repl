package main

import (
	"fmt"
)

func init() {
	Register("help", "Displays a help message", Help)
}

// Help displays all available commands and their descriptions.
// It iterates through the registered Commands map and prints each command's
// name and description in a formatted list.
func Help(config *Config, args []string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n")

	for _, command := range CommandsList {
		fmt.Printf("   - %s: %s\n", command.Name(), command.Description())
	}
	return nil
}
