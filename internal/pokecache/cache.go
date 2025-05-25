package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	Entries  map[string]cacheEntry
	Mutex    sync.Mutex
	interval time.Duration
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	done := make(chan bool)
	for {
		select {
		case <-done:
			fmt.Println("reap loop done")
			return
		case <-ticker.C:
			for k, v := range c.Entries {
				if time.Since(v.createdAt) > c.interval {
					delete(c.Entries, k)
				}
			}
		}
	}
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		Entries:  make(map[string]cacheEntry, 0),
		Mutex:    sync.Mutex{},
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	entry, ok := c.Entries[key]
	return entry.val, ok
}
