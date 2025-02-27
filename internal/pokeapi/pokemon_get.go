package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
)

func (c *Client) GetPokemon(name string) (bool, Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	var data []byte

	// check if resp is already cached
	// if cached response is found use local data
	if entry, exists := c.cache.Get(url); exists {
		data = entry
	} else { // else make call to pokeapi
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return false, Pokemon{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return false, Pokemon{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return false, Pokemon{}, fmt.Errorf("%s", res.Status)
		}

		data, err = io.ReadAll((res.Body))
		if err != nil {
			return false, Pokemon{}, err
		}

		c.cache.Add(url, data) // add result to the cache
	}

	var pokemonResp Pokemon
	err := json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return false, Pokemon{}, err
	}

	if tryCatch(pokemonResp.BaseExperience) {
		return true, pokemonResp, nil
	}

	return false, Pokemon{}, nil

}

func tryCatch(baseExperience int) bool {
	// Calculate the probability for true based on the base value
	var probability float64

	// For base >= 340, cap probability at 25% (0.25)
	if baseExperience >= 340 {
		probability = 0.25
	} else {
		// Linearly decrease probability from 99% at base 36 to 25% at base 340
		probability = 0.99 - (0.74 * float64(baseExperience-36) / float64(340-36))
	}

	// Generate a random number between 0.0 and 1.0
	randomValue := rand.Float64()

	// Return true if the random value is less than the probability
	return randomValue < probability
}
