package pokeapi_test

import (
	"testing"
	"time"

	"github.com/weiwenchen2022/pokedexcli/internal/pokeapi"
)

func TestLocationAreaByName(t *testing.T) {
	t.Parallel()

	c := pokeapi.NewConfig(5 * time.Second)

	id := "pastoria-city-area"
	result, err := pokeapi.LocationArea(c, id)
	if err != nil {
		t.Error(err)
	}

	if id != result.Name {
		t.Errorf("Expect receive %s area.", id)
	}
}
