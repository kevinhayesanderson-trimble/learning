package main

import (
	"fmt"
	"sync"
)

type SafeCache struct {
	mu      sync.Mutex
	visited map[string]bool
}

func (cache *SafeCache) IsVisitedAndMarked(url string) bool {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if cache.visited[url] {
		return true
	}

	cache.visited[url] = true
	return false
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *SafeCache, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth <= 0 {
		return
	}
	if cache.IsVisitedAndMarked(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go func(url string) {
			Crawl(url, depth-1, fetcher, cache, wg)
		}(u)
	}
}

// Strategy 1 (Shared Memory + Mutex + WaitGroup)
func test1() {
	cache := &SafeCache{visited: make(map[string]bool)}
	var wg sync.WaitGroup
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, cache, &wg)
	wg.Wait()
}
