package InMemoryCache

// server/pkg/in-memory-cache

import (
	"sync"
	"time"
)

type Item[V any] struct{
	value V
	expiry time.Time 
}

type Cache[K comparable, V any] struct {
	items map[K]Item[V]
	mu sync.Mutex
}

func (i Item[V]) isExpired() bool {
	return time.Now().After(i.expiry)
}

// new cache instance
func New[K comparable, V any]() *Cache[K,V]{
	c := &Cache[K, V]{
		items: make(map[K]Item[V]),
	}

	go func() {
		for range time.Tick(5 * time.Second){
			c.mu.Lock()

			for key, item := range c.items {
				if item.isExpired() {
					delete(c.items, key)
				}
			}

			c.mu.Unlock()
		}
	}()

	return c
}

// This method has a receiver type 
// assigns Set to type Cache similar to classes in OOP
func (c *Cache[K,V]) Set(key K, value V, ttl time.Duration){
	c.mu.Lock()
	// defers this command until Set() is finished
	defer c.mu.Unlock()

	c.items[key] = Item[V]{
		value: value,
		expiry: time.Now().Add(ttl),
	}
}

// gets a value from the cache
// the (V, bool) is the return type
func (c *Cache[K,V]) Get(key K) (V, bool){
	c.mu.Lock()
	defer c.mu.Unlock()

	item, found := c.items[key]

	if !found {
		return item.value, false
	}

	if item.isExpired() {
		delete(c.items, key)
	}

	return item.value, found
}

func (c *Cache[K,V]) Remove(key K){
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items,key)
}

// pops item from map and returns to caller
func (c *Cache[K,V]) Pop(key K)(V, bool){
	c.mu.Lock()
	defer c.mu.Unlock()

	item, found := c.items[key]

	if !found {
		return item.value, false
	}

	delete(c.items,key)

	return item.value, found
}

