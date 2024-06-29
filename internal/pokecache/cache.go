package pokecache

import (
	"sync"
	"time"
)

// Creates a new Cache struct to be used by the commands.
// This is intended to be called in the pokiapi.Client struct along with the http.Client in order to give the http client a cache.
func NewCache(ttl time.Duration) Cache {
	cache := Cache{
		// create new cache map
		cache: make(map[string]cacheEntry),
		// create a new mutex for concurrent map serialization
		mu: &sync.RWMutex{},
	}

	go cache.reapLoop(ttl)

	return cache
}

// It's important to use this method when adding entries to the cache because it imposes a mutex and sets the time of entry.
// This is important to the `reapLoop()` function in order to free up cache space after a certain period of time.
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		Val:       val,
	}
}

// It's important to use this method when reading cache entries because it imposes a mutex on the map.
// This is important because the `reapLoop()` function is called on a different goroutine, and is continuously running along-side the main goroutine.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.cache[key]
	if !ok {
		return []byte{}, false
	}

	return entry.Val, ok
}

func (c *Cache) reapLoop(ttl time.Duration) {
	ticker := time.NewTicker(ttl)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()

		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > ttl {
				delete(c.cache, key)
			}
		}

		c.mu.Unlock()
	}
}
