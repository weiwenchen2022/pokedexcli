package pokeapi

import (
	"encoding/json"
	"errors"
	"math/rand"
	"os"
	"time"

	"github.com/weiwenchen2022/pokedexcli/internal/pokeapi/structs"
	"github.com/weiwenchen2022/pokedexcli/internal/pokecache"
)

type Pokedex map[string]structs.Pokemon

type Config struct {
	next     string
	previous string

	cache *pokecache.Cache

	r *rand.Rand

	Pokedex Pokedex
}

func NewConfig(interval time.Duration) *Config {
	return &Config{
		next: locationAreaEndpoint,

		cache: pokecache.NewCache(interval),

		r:       rand.New(rand.NewSource(time.Now().UnixNano())),
		Pokedex: make(Pokedex),
	}
}

func (c *Config) Load(filepath string) error {
	f, err := os.Open(filepath)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	} else if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&c.Pokedex)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) Save(filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")

	if err := enc.Encode(c.Pokedex); err != nil {
		return err
	}

	if err := f.Sync(); err != nil {
		return err
	}

	return nil
}
