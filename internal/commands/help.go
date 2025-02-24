package commands

import (
	"fmt"
)

func init() {
	Register("help", "Displays a help message", Help)
}

func Help([]string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")

	for _, command := range Commands {
		fmt.Printf("%s: %s\n", command.Name(), command.Description())
	}
	return nil
}
