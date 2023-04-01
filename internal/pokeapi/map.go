package pokeapi

import (
	"bytes"
	"errors"
)

type pokeAPIAreaResponse struct {
	Count    int `json:"count"`
	Next     any `json:"next"`
	Previous any `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const locationAreaEndpoint = "https://pokeapi.co/api/v2/location-area/"

func Next(c *Config) ([]string, error) {
	if c.next == "" {
		return nil, errors.New("no next location areas")
	}

	return locations(c, c.next)
}

func Previous(c *Config) ([]string, error) {
	if c.previous == "" {
		return nil, errors.New("no invoke map command twice before")
	}

	return locations(c, c.previous)
}

func locations(c *Config, endpoint string) ([]string, error) {
	body, err := makeRequest(c, endpoint)
	if err != nil {
		return nil, err
	}

	var response pokeAPIAreaResponse
	err = parseResponse(bytes.NewReader(body), &response)
	if err != nil {
		return nil, err
	}

	names := make([]string, len(response.Results))
	for i := range names {
		names[i] = response.Results[i].Name
	}

	if next, ok := response.Next.(string); ok {
		c.next = next
	} else {
		c.next = ""
	}

	if previous, ok := response.Previous.(string); ok {
		c.previous = previous
	} else {
		c.previous = ""
	}

	return names, nil
}
