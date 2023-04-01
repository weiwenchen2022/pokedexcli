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
	mu       sync.Mutex
	data     map[string]cacheEntry
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		data:     make(map[string]cacheEntry),
		interval: interval,
	}

	go c.reaploop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.data[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, ok := c.data[key]; ok {
		return entry.val, true
	}

	return nil, false
}

func (c *Cache) reaploop() {
	ticker := time.NewTicker(c.interval)

	for {
		<-ticker.C

		now := time.Now()

		c.mu.Lock()
		for k, v := range c.data {
			if now.Sub(v.createdAt) > c.interval {
				delete(c.data, k)
			}
		}
		c.mu.Unlock()
	}
}
