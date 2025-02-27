package main

import (
	"fmt"
)

func init() {
	Register("catch", "catch takes a pokemon name and attempts to catch that pokemon with a pokeball", Catch)
}

func Catch(cfg *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	caught, pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	if caught {
		fmt.Printf("%s was caught!\n", name)
		cfg.caughtPokemon[name] = pokemon
	} else {
		fmt.Printf("%s excaped!\n", name)
	}
	return nil
}
