package goquest

import (
	"sync"
	"time"
)

type cache[T any] struct {
	values map[string]*cacheEntry[T]
	mu     sync.Mutex
}

type cacheEntry[T any] struct {
	value     any
	expiresAt time.Time
}

func newCache[T any]() *cache[T] {
	return &cache[T]{
		values: make(map[string]*cacheEntry[T]),
	}
}
