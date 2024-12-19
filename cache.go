package main
import (
	"errors"
	"sync"
)
var ErrConditionNotFound = errors.New("condition not found")

// Cache structure to store query results
type Cache struct {
	mu    sync.RWMutex
	store map[string][]SquirrelData
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string][]SquirrelData),
	}
}

func (c *Cache) Get(key string) ([]SquirrelData, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	data, exists := c.store[key]
	return data, exists
}

func (c *Cache) Set(key string, data []SquirrelData) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = data
}