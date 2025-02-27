package main

import "fmt"

func init() {
	Register("inspect", "display details about a specified pokemon", Inspect)
}

func Inspect(cfg *Config, name []string) error {
	pokemon, ok := cfg.caughtPokemon[name[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("   -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, typing := range pokemon.Types {
		fmt.Printf("   -%s\n", typing.Type.Name)
	}

	return nil
}
