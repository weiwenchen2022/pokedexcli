package pokeapi

const exploreEndpoint = "https://pokeapi.co/api/v2/location-area/%s/"

func Explore(c *Config, id string) ([]string, error) {
	l, err := LocationArea(c, id)
	if err != nil {
		return nil, err
	}

	names := make([]string, len(l.PokemonEncounters))
	for i := range names {
		names[i] = l.PokemonEncounters[i].Pokemon.Name
	}

	return names, nil
}
