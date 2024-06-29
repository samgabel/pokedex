package pokeapi

import (
	"net/http"

	"github.com/samgabel/pokedex/internal/pokecache"
)

// This struct handles our requests as well as our caching of said requests
type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

type locationObj struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`     // nullable string (string or nil pointer)
	Previous *string `json:"previous"` // nullable string (string or nil pointer)
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
