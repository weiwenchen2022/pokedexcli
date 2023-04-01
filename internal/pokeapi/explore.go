package pokeapi

import (
	"bytes"
	"fmt"
)

type pokeAPIExploreResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

const exploreEndpoint = "https://pokeapi.co/api/v2/location-area/%s/"

func Explore(c *Config, location string) ([]string, error) {
	endpoint := fmt.Sprintf(exploreEndpoint, location)

	body, err := makeRequest(c, endpoint)
	if err != nil {
		return nil, err
	}

	var response pokeAPIExploreResponse
	err = parseResponse(bytes.NewReader(body), &response)
	if err != nil {
		return nil, err
	}

	names := make([]string, len(response.PokemonEncounters))
	for i := range names {
		names[i] = response.PokemonEncounters[i].Pokemon.Name
	}

	return names, nil
}
