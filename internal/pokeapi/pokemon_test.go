package pokeapi_test

import (
	"testing"
	"time"

	"github.com/weiwenchen2022/pokedexcli/internal/pokeapi"
)

func TestPokemonByName(t *testing.T) {
	t.Parallel()

	c := pokeapi.NewConfig(1 * time.Second)

	id := "pikachu"
	result, err := pokeapi.Pokemon(c, id)
	if err != nil {
		t.Error(err)
	}

	if id != result.Name {
		t.Errorf("Expect to receive %s", id)
	}
}
