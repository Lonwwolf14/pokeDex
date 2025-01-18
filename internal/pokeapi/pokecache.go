package pokeapi

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	store    map[string]cacheEntry
	interval time.Duration
	mu       sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	return &Cache{
		store:    make(map[string]cacheEntry),
		interval: interval,
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.store[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	defer c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.store[key]
	if !ok {
		return nil, false
	}
	if c.isExpired(entry) {
		delete(c.store, key)
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) isExpired(entry cacheEntry) bool {
	return time.Since(entry.createdAt) > c.interval
}

func (c *Cache) repLoop() {
	c.mu.Lock()
	for key, entry := range c.store {
		if c.isExpired(entry) {
			delete(c.store, key)
		}
	}
	defer c.mu.Unlock()
}

func (c *Cache) StartCleanup() {
	go func() {
		ticker := time.NewTicker(c.interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				c.repLoop()
			}
		}
	}()
}
