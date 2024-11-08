package goquest

import (
	"context"
	"sync"
	"time"
)

type cache[T any] struct {
	values map[string]*cacheEntry[T]
	mu     sync.Mutex
}

type cacheEntry[T any] struct {
	value     *FetchResult[T]
	expiresAt time.Time
}

func newCache[T any]() *cache[T] {
	return &cache[T]{
		values: make(map[string]*cacheEntry[T]),
	}
}

// fetch retrieves a value from the cache or uses fetchFunc to get it and store it if not cached or expired.
func (c *cache[T]) fetch(
	ctx context.Context,
	key string,
	fetchFunc func(ctx context.Context) (*FetchResult[T], error),
	cacheExpiration time.Duration,
) (*FetchResult[T], error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Check if the cache has a valid entry for the given key.
	if entry, ok := c.values[key]; ok && entry.expiresAt.After(time.Now()) {
		return entry.value, nil
	}

	// Fetch the result using the provided fetchFunc.
	result, err := fetchFunc(ctx)
	if err != nil {
		return nil, err
	}

	// Store the fetched result in the cache with expiration time.
	c.values[key] = &cacheEntry[T]{
		value:     result,
		expiresAt: time.Now().Add(cacheExpiration),
	}

	return result, nil
}
