package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	go c.repLoop(interval)
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	defer c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) repLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for key, entry := range c.cache {
		if entry.createdAt.Before(now.Add(-last)) {
			delete(c.cache, key)
		}
	}
}
