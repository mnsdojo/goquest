// middleware.go
package goquest

import (
	"context"
	"net/http"
)

type Middleware[T any] func(ctx context.Context, req *http.Request, result *FetchResult[T]) (context.Context, *http.Request, *FetchResult[T])

