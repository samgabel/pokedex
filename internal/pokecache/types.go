package pokecache

import (
	"sync"
	"time"
)

// This struct is to be used as our map value for cache in the Cache struct.
type cacheEntry struct {
	createdAt time.Time
	Val       []byte
}

// The Cache struct in pokecache serves to cache our API responses, reducing the number of our API calls.
type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.RWMutex
}
