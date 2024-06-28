package pokeapi

import "net/http"

type Client struct {
	httpClient http.Client
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
