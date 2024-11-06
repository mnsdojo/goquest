package goquest

import (
	"io"
	"time"
)

// HTTPMethod represents the HTTP method to be used for a fetch request.
type HTTPMethod string

const (
	Get    HTTPMethod = "GET"
	Post   HTTPMethod = "POST"
	Put    HTTPMethod = "PUT"
	Delete HTTPMethod = "DELETE"
	Patch  HTTPMethod = "PATCH" // Added PATCH method
)

// fetchoptions <-> contains the configuration for a fetch operation.
type FetchOptions struct {
	Body            io.Reader
	Headers         map[string]string
	Method          HTTPMethod
	CacheExpiration time.Duration
}

// FetchResult represents the result of a fetch operation.
type FetchResult[T any] struct {
	Data   T
	Error  error
	Status int
}
