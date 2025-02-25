package commands

// Map functions handle the Pokemon location area exploration functionality.
// It provides commands for navigating through paginated lists of Pokemon locations.

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func init() {
	Register("map", "Display next 20 pokemon locations", Map)
	Register("mapb", "Display the previous 20 pokemon locations", Mapb)
}

// Map displays the next page of location areas, using the Next URL from config
// or the base URL if starting fresh.
func Map(config *Config, args []string) error {
	var url string
	if config.Next != nil {
		url = *config.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area"
	}
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to get map data")
		return err
	}
	defer res.Body.Close()

	locationAreas, err := printLocations(res.Body)
	if err != nil {
		return err
	}

	updateConfig(config, locationAreas.Next, locationAreas.Previous)
	return nil
}

// Mapb displays the previous page of location areas if available.
// Returns an error if no previous page exists.
func Mapb(config *Config, args []string) error {
	if config.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(*config.Previous)
	if err != nil {
		fmt.Println("Error retrieving previous locations")
		return err
	}
	defer res.Body.Close()

	locationAreas, err := printLocations(res.Body)
	if err != nil {
		return err
	}

	updateConfig(config, locationAreas.Next, locationAreas.Previous)
	return nil
}

// printLocations parses the API response body and prints location names.
// Returns the parsed LocationAreaList and any errors encountered.
func printLocations(body io.ReadCloser) (LocationAreaList, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return LocationAreaList{}, err
	}

	var locationAreas LocationAreaList
	if err = json.Unmarshal(data, &locationAreas); err != nil {
		fmt.Println("Unable to format data")
		return LocationAreaList{}, err
	}

	for i := 0; i < len(locationAreas.Results); i++ {
		fmt.Println(locationAreas.Results[i].Name)
	}

	return locationAreas, nil
}

// updateConfig updates the pagination URLs in the config struct.
func updateConfig(config *Config, next *string, previous *string) {
	config.Next = next
	config.Previous = previous
}

// Result represents a single location area entry from the PokeAPI.
type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// LocationAreaList represents the paginated response from the PokeAPI location-area endpoint.
type LocationAreaList struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`     // Nullable URL for next page
	Previous *string  `json:"previous"` // Nullable URL for previous page
	Results  []Result `json:"results"`
}
