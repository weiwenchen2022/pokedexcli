package pokeapi

const pokemonEndpoint = "https://pokeapi.co/api/v2/pokemon/%s/"

func Catch(c *Config, name string) (bool, error) {
	p, err := Pokemon(c, name)
	if err != nil {
		return false, err
	}

	if c.r.Intn(100)+100 < p.BaseExperience {
		return false, nil
	}

	c.Pokedex[name] = p
	return true, nil
}
