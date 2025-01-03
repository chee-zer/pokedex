package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	Val []byte
}

type Cache struct{
	Entry map[string]cacheEntry
	interval time.Duration
	mu sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {

	c := &Cache {
		Entry: make(map[string]cacheEntry),
		interval: interval,
		mu: sync.RWMutex{},

	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add (key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry {
		createdAt: time.Now(),
		Val: val,
	}
	c.Entry[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, exists := c.Entry[key]
	if exists {
		return entry.Val, true
	} else {
		return []byte{}, false
	}
}

func (c * Cache) delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.Entry, key)
}

func (c *Cache) reapLoop() {
	timer := time.NewTicker(c.interval)
	for {
		<- timer.C
		now := time.Now()
		for k, v := range c.Entry {
			age := v.createdAt.Add(c.interval)
			if age.Before(now) {
				c.delete(k)
			}
		}
	}
}