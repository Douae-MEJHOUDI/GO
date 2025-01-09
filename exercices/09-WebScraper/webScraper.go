package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	c := context.Background()

	urls := []string{"http://example.com", "http://httpbin.org/delay/5", "http://httpbin.org/delay/1"}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			ctx, cancel := context.WithTimeout(c, time.Second*2)
			defer cancel()
			defer wg.Done()
			fetchURL(ctx, url)

		}(url)
	}

	wg.Wait()

}

func fetchURL(ctx context.Context, url string) {
	done := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		done <- url
	}()
	select {
	case <-ctx.Done():
		fmt.Println("URL fetching cancelled:", ctx.Err())
	case <-done:
		fmt.Println("Fetching URL:", url)
	}
}
