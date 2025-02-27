package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListEncounters(id string) (RespAreaEncounters, error) {
	url := baseURL + "/location-area/" + id
	var data []byte

	// check if resp is already cached
	// if cached response is found use local data
	if entry, exists := c.cache.Get(url); exists {
		data = entry
	} else { // else make call to pokeapi
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespAreaEncounters{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespAreaEncounters{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return RespAreaEncounters{}, fmt.Errorf("%s", res.Status)
		}

		data, err = io.ReadAll((res.Body))
		if err != nil {
			return RespAreaEncounters{}, err
		}

		c.cache.Add(url, data) // add result to the cache
	}

	var exploreResp RespAreaEncounters
	err := json.Unmarshal(data, &exploreResp)
	if err != nil {
		return RespAreaEncounters{}, err
	}

	return exploreResp, nil
}
