package pokeapi

import (
	"net/http"
	"time"

	"github.com/samgabel/pokedex/internal/pokecache"
)

// Creates a new http.Client and pokecache.Cache.
// This way we can group together the requester and the cache for the response bodies.
// We want to specify the http client request timeout (probably in seconds) and the cache time-to-live for each entry (probably in minutes).
// Meant to be called early (and once per REPL session) in the program by `startRepl()`
func NewClient(httpTimeout, cacheTTL time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: httpTimeout,
		},
		cache: pokecache.NewCache(cacheTTL),
	}
}

