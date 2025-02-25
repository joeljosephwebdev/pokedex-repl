package main

// Map functions handle the Pokemon location area exploration functionality.
// It provides commands for navigating through paginated lists of Pokemon locations.

import (
	"errors"
	"fmt"

	"github.com/joeljosephwebdev/pokedex-repl/internal/pokeapi"
)

func init() {
	Register("map", "Display next 20 pokemon locations", Map)
	Register("mapb", "Display the previous 20 pokemon locations", Mapb)
}

// Map displays the next page of location areas, using the Next URL from config
// or the base URL if starting fresh.
func Map(cfg *Config, args []string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	printLocations(locationsResp)
	updateConfig(cfg, locationsResp.Next, locationsResp.Previous)

	return nil
}

// Mapb displays the previous page of location areas if available.
// Returns an error if no previous page exists.
func Mapb(cfg *Config, args []string) error {
	if cfg.previousLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)
	if err != nil {
		return err
	}

	printLocations(locationsResp)
	updateConfig(cfg, locationsResp.Next, locationsResp.Previous)

	return nil
}

// printLocations parses the API response body and prints location names.
// Returns the parsed LocationAreaList and any errors encountered.
func printLocations(locationAreas pokeapi.RespShallowLocations) {
	for i := 0; i < len(locationAreas.Results); i++ {
		fmt.Println(locationAreas.Results[i].Name)
	}
}

// updateConfig updates the pagination URLs in the config struct.
func updateConfig(cfg *Config, next *string, previous *string) {
	cfg.nextLocationsURL = next
	cfg.previousLocationsURL = previous
}
