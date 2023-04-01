package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"pokedexcli/internal/pokecache"
)

type Pokedex map[string]Pokemon

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

func makeRequest(c *Config, endpoint string) ([]byte, error) {
	if body, ok := c.cache.Get(endpoint); ok {
		return body, nil
	}

	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	c.cache.Add(endpoint, body)
	return body, nil
}

func parseResponse(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(v)
}
