package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/mnsdojo/goquest"
)

type Post struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

func main() {

	client := &http.Client{}

	fetcher := goquest.NewFetcher[Post](client, 5*time.Minute)
	ctx := context.Background()
	getResult, err := fetcher.Get(ctx, "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("Error performing GET request:", err)
		return
	}
	fmt.Printf("GET Response: %+v\n", getResult.Data)

}
