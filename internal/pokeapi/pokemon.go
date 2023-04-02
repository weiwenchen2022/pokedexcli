package pokeapi

import (
	"fmt"

	"github.com/weiwenchen2022/pokedexcli/internal/pokeapi/structs"
)

func Pokemon(c *Config, id string) (result structs.Pokemon, err error) {
	err = do(c, fmt.Sprintf("pokemon/%s", id), &result)
	return
}
