package pokeapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const apiurl = "https://pokeapi.co/api/v2/"

var client = &http.Client{}

func do(c *Config, endpoint string, result any) error {
	if body, ok := c.cache.Get(endpoint); ok {
		return json.NewDecoder(bytes.NewReader(body)).Decode(result)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiurl+endpoint, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	if resp.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n",
			resp.StatusCode, body)
	}

	if err != nil {
		return err
	}

	c.cache.Add(endpoint, body)
	return json.NewDecoder(bytes.NewReader(body)).Decode(result)
}
