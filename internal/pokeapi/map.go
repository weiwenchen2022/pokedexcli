package pokeapi

import (
	"errors"
	"strings"
)

const locationAreaEndpoint = "location-area/"

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
	l, err := Resource(c, endpoint)
	if err != nil {
		return nil, err
	}

	names := make([]string, len(l.Results))
	for i := range names {
		names[i] = l.Results[i].Name
	}

	if next, ok := l.Next.(string); ok {
		index := strings.Index(next, locationAreaEndpoint) + len(locationAreaEndpoint)
		c.next = locationAreaEndpoint + next[index:]
	} else {
		c.next = ""
	}

	if previous, ok := l.Previous.(string); ok {
		index := strings.Index(previous, locationAreaEndpoint) + len(locationAreaEndpoint)
		c.previous = locationAreaEndpoint + previous[index:]
	} else {
		c.previous = ""
	}

	return names, nil
}
