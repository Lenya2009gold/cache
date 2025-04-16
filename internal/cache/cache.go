package cache

import (
	"container/list"
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]*list.Element
	order    *list.List
	capacity int
	mu       sync.RWMutex
	TTL      time.Duration
}

func New(capacity int, ttl time.Duration) *Cache {
	return &Cache{
		cache:    make(map[string]*list.Element),
		order:    list.New(),
		capacity: capacity,
		TTL:      ttl,
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, exists := c.cache[key]
	if exists {
		c.order.MoveToFront(item)
		return
	}
	if c.order.Len() >= c.capacity {
		c.removeOldest()
	}
	c.cache[key] = c.order.PushFront(value)

}
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, exists := c.cache[key]
	if !exists {
		return nil, false
	}
	return item.Value, true
}
func (c *Cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	element, exists := c.cache[key]
	if !exists {
		return

	}
	c.order.Remove(element)
	delete(c.cache, key)
}

func (c *Cache) removeOldest() {
	if c.order.Len() == 0 {
		return
	}
	oldest := c.order.Back()
	c.order.Remove(oldest)
	delete(c.cache, oldest.Value.(string))
}
