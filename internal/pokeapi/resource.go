package pokeapi

import "github.com/weiwenchen2022/pokedexcli/internal/pokeapi/structs"

func Resource(c *Config, endpoint string) (result structs.Resource, err error) {
	err = do(c, endpoint, &result)
	return
}
