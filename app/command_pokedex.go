package main

import (
	"fmt"
)

func init() {
	Register("pokedex", "See all of your caught pokemon", Pokedex)
}

func Pokedex(cfg *Config, args []string) error {
	if len(cfg.caughtPokemon) < 1 {
		fmt.Println("You have not caught any pokemon yet.\nYou may try with the catch command.")
		return nil
	}

	fmt.Println("Your Pokedex: ")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("   - %s\n", pokemon.Name)
	}
	return nil
}
