package main

import (
	"fmt"
)

func init() {
	Register("explore", "Displays all pokemon that can be found at provided location.", Explore)
}

// Explore takes an area name or id and displays a full list of pokemon found in that area
func Explore(cfg *Config, args []string) error {
	exploreResp, err := cfg.pokeapiClient.ListEncounters(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", exploreResp.Location.Name)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range exploreResp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
