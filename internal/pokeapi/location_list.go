package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area/?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	var data []byte

	// check if resp is already cached
	// if cached response is found use local data
	if entry, exists := c.cache.Get(url); exists {
		data = entry
	} else { // else make call to pokeapi
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocations{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocations{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll((res.Body))
		if err != nil {
			return RespShallowLocations{}, err
		}
		c.cache.Add(url, data) // add result to the cache
	}

	var locationsResp RespShallowLocations
	err := json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
