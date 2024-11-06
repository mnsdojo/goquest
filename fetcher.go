package goquest

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type Fetcher[T any] struct {
	cache             *cache[T]
	client            *http.Client
	defaultExpiration time.Duration // Default cache expiration for all fetch operations
}

//

func NewFetcher[T any](httpClient *http.Client, defaultExpiration time.Duration) *Fetcher[T] {
	if httpClient == nil {
		httpClient = http.DefaultClient // Use the default HTTP client if none is provided
	}

	return &Fetcher[T]{
		cache:             newCache[T](),
		client:            httpClient,
		defaultExpiration: defaultExpiration,
	}
}

// get reqwst with default expiration
func (f *Fetcher[T]) Get(ctx context.Context, url string, headers map[string]string) (*FetchResult[T], error) {
	return f.Fetch(ctx, url, FetchOptions{
		Method:          Get,
		CacheExpiration: f.defaultExpiration, // Use the default expiration
		Headers:         headers,
	})
}

func (f *Fetcher[T]) Fetch(ctx context.Context, url string, options FetchOptions) (*FetchResult[T], error) {
	// Define a function to do  the HTTP request
	fetchFunc := func(ctx context.Context) (*FetchResult[T], error) {
		req, err := http.NewRequest(string(options.Method), url, options.Body)
		if err != nil {
			return nil, err
		}

		// Set headers
		for key, value := range options.Headers {
			req.Header.Set(key, value)
		}

		// Perform the request
		resp, err := f.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		// Parse response and return the result
		var result T
		// Assuming the response body can be unmarshalled directly into `T`
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, err
		}

		return &FetchResult[T]{Data: result, Status: resp.StatusCode}, nil
	}

	// Use the cache to fetch the result
	return f.cache.fetch(ctx, url, fetchFunc, options.CacheExpiration)
}
//