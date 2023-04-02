package pokeapi

import (
	"fmt"

	"github.com/weiwenchen2022/pokedexcli/internal/pokeapi/structs"
)

func LocationArea(c *Config, id string) (result structs.LocationArea, err error) {
	err = do(c, fmt.Sprintf("location-area/%s", id), &result)
	return
}
