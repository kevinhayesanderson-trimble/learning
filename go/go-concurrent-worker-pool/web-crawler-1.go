package main

import "fmt"

type CrawlResult struct {
	depth int
	url   string
	body  string
	urls  []string
	err   error
}

func Worker(url string, depth int, fetcher Fetcher, out chan CrawlResult) {
	body, urls, err := fetcher.Fetch(url)
	out <- CrawlResult{
		depth: depth,
		url:   url,
		body:  body,
		urls:  urls,
		err:   err,
	}
}

func WorkerWithSignal(url string, depth int, fetcher Fetcher, dataChan chan CrawlResult, signalChan chan struct{}) {
	defer func() { signalChan <- struct{}{} }()
	body, urls, err := fetcher.Fetch(url)
	result := CrawlResult{
		depth: depth,
		url:   url,
		body:  body,
		urls:  urls,
		err:   err,
	}
	select {
	case dataChan <- result:
	default:
		fmt.Printf("System slammed. Dropped data for: %s\n", url)
	}
}

func Crawl1(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	dataChan := make(chan CrawlResult, 10)
	signalChan := make(chan struct{})

	visited := make(map[string]bool)
	workerCount := 0

	visited[url] = true
	workerCount++
	go WorkerWithSignal(url, depth, fetcher, dataChan, signalChan)

	for workerCount > 0 {
		select {
		case res := <-dataChan:

			if res.err != nil {
				fmt.Println(res.err)
				continue
			}
			fmt.Printf("found: %s %q\n", res.url, res.body)

			if res.depth > 0 {
				for _, resurl := range res.urls {
					if !visited[resurl] {
						visited[resurl] = true
						workerCount++
						go WorkerWithSignal(resurl, res.depth-1, fetcher, dataChan, signalChan)
					}
				}
			}
		case <-signalChan:
			workerCount--
		}
	}
}

// Strategy 2: Channel Coordinator
func test2() {
	Crawl1("https://golang.org/", 4, fetcher)
}
