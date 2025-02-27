package main

import (
	"fmt"
)

func init() {
	Register("explore", "Displays all pokemon that can be found at the provided location", Explore)
}

// Explore takes an area name or id and displays a full list of pokemon found in that area
func Explore(cfg *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("you must provide a location name or id")
	}

	exploreResp, err := cfg.pokeapiClient.GetLocation(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", exploreResp.Location.Name)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range exploreResp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
