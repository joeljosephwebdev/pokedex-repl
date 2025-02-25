package main

// Exit provides the exit functionality to close the Pokedex application.

import (
	"fmt"
	"os"
)

func init() {
	Register("exit", "Exits the pokedex", Exit)
}

// Exit gracefully terminates the Pokedex application.
// Note: The return statement is never reached due to os.Exit(0).
func Exit(config *Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
