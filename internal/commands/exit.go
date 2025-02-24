package commands

import (
	"fmt"
	"os"
)

func init() {
	Register("exit", "Exits the pokedex", Exit)
}

func Exit([]string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
