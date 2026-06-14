package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// CacheItem represents a cached item with expiration
type CacheItem struct {
	Value      interface{}
	Expiration time.Time
}

// InMemoryCache implements an in-memory cache (fallback when Redis is not available)
type InMemoryCache struct {
	items map[string]CacheItem
	mu    sync.RWMutex
}

// NewInMemoryCache creates a new in-memory cache instance
func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		items: make(map[string]CacheItem),
	}
}

// NewRedisCache creates a new cache instance
// For now, returns an in-memory cache as fallback
// TODO: Implement actual Redis connection when dependency is added
func NewRedisCache(addr, password string, db int) (*InMemoryCache, error) {
	return NewInMemoryCache(), nil
}

// Close closes the cache connection
func (imc *InMemoryCache) Close() error {
	imc.mu.Lock()
	defer imc.mu.Unlock()
	imc.items = make(map[string]CacheItem)
	return nil
}

// Set stores a value in the cache
func (imc *InMemoryCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	imc.mu.Lock()
	defer imc.mu.Unlock()

	var exp time.Time
	if expiration > 0 {
		exp = time.Now().Add(expiration)
	}

	imc.items[key] = CacheItem{
		Value:      value,
		Expiration: exp,
	}

	return nil
}

// Get retrieves a value from the cache
func (imc *InMemoryCache) Get(ctx context.Context, key string, dest interface{}) error {
	imc.mu.RLock()
	defer imc.mu.RUnlock()

	item, exists := imc.items[key]
	if !exists {
		return fmt.Errorf("key not found")
	}

	if !item.Expiration.IsZero() && time.Now().After(item.Expiration) {
		return fmt.Errorf("key not found")
	}

	jsonData, err := json.Marshal(item.Value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	if err := json.Unmarshal(jsonData, dest); err != nil {
		return fmt.Errorf("failed to unmarshal value: %w", err)
	}

	return nil
}

// Delete removes a value from the cache
func (imc *InMemoryCache) Delete(ctx context.Context, key string) error {
	imc.mu.Lock()
	defer imc.mu.Unlock()
	delete(imc.items, key)
	return nil
}

// DeletePattern removes all keys matching a pattern (simple prefix match)
func (imc *InMemoryCache) DeletePattern(ctx context.Context, pattern string) error {
	imc.mu.Lock()
	defer imc.mu.Unlock()

	for key := range imc.items {
		if len(key) >= len(pattern) && key[:len(pattern)] == pattern {
			delete(imc.items, key)
		}
	}
	return nil
}

// Exists checks if a key exists in the cache
func (imc *InMemoryCache) Exists(ctx context.Context, key string) (bool, error) {
	imc.mu.RLock()
	defer imc.mu.RUnlock()

	item, exists := imc.items[key]
	if !exists {
		return false, nil
	}

	if !item.Expiration.IsZero() && time.Now().After(item.Expiration) {
		return false, nil
	}

	return true, nil
}

// SetMultiple stores multiple values in the cache
func (imc *InMemoryCache) SetMultiple(ctx context.Context, items map[string]interface{}, expiration time.Duration) error {
	imc.mu.Lock()
	defer imc.mu.Unlock()

	var exp time.Time
	if expiration > 0 {
		exp = time.Now().Add(expiration)
	}

	for key, value := range items {
		imc.items[key] = CacheItem{
			Value:      value,
			Expiration: exp,
		}
	}

	return nil
}

// GetMultiple retrieves multiple values from the cache
func (imc *InMemoryCache) GetMultiple(ctx context.Context, keys []string) (map[string]interface{}, error) {
	imc.mu.RLock()
	defer imc.mu.RUnlock()

	result := make(map[string]interface{})
	now := time.Now()

	for _, key := range keys {
		item, exists := imc.items[key]
		if !exists {
			continue
		}

		if !item.Expiration.IsZero() && now.After(item.Expiration) {
			continue
		}

		result[key] = item.Value
	}

	return result, nil
}

// Increment increments a numeric value in the cache
func (imc *InMemoryCache) Increment(ctx context.Context, key string) (int64, error) {
	imc.mu.Lock()
	defer imc.mu.Unlock()

	item, exists := imc.items[key]
	var value int64 = 1

	if exists {
		if v, ok := item.Value.(int64); ok {
			value = v + 1
		} else if v, ok := item.Value.(int); ok {
			value = int64(v) + 1
		}
	}

	imc.items[key] = CacheItem{
		Value:      value,
		Expiration: item.Expiration,
	}

	return value, nil
}

// Decrement decrements a numeric value in the cache
func (imc *InMemoryCache) Decrement(ctx context.Context, key string) (int64, error) {
	imc.mu.Lock()
	defer imc.mu.Unlock()

	item, exists := imc.items[key]
	var value int64 = -1

	if exists {
		if v, ok := item.Value.(int64); ok {
			value = v - 1
		} else if v, ok := item.Value.(int); ok {
			value = int64(v) - 1
		}
	}

	imc.items[key] = CacheItem{
		Value:      value,
		Expiration: item.Expiration,
	}

	return value, nil
}

// Expire sets an expiration time for a key
func (imc *InMemoryCache) Expire(ctx context.Context, key string, expiration time.Duration) error {
	imc.mu.Lock()
	defer imc.mu.Unlock()

	item, exists := imc.items[key]
	if !exists {
		return fmt.Errorf("key not found")
	}

	imc.items[key] = CacheItem{
		Value:      item.Value,
		Expiration: time.Now().Add(expiration),
	}

	return nil
}

// TTL returns the remaining time to live of a key
func (imc *InMemoryCache) TTL(ctx context.Context, key string) (time.Duration, error) {
	imc.mu.RLock()
	defer imc.mu.RUnlock()

	item, exists := imc.items[key]
	if !exists {
		return 0, fmt.Errorf("key not found")
	}

	if item.Expiration.IsZero() {
		return -1, nil
	}

	ttl := time.Until(item.Expiration)
	if ttl < 0 {
		return 0, nil
	}

	return ttl, nil
}

// FlushAll removes all keys from the cache
func (imc *InMemoryCache) FlushAll(ctx context.Context) error {
	imc.mu.Lock()
	defer imc.mu.Unlock()
	imc.items = make(map[string]CacheItem)
	return nil
}
