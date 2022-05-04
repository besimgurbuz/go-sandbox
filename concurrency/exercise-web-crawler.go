package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page
	Fetch(url string) (body string, urls []string, err error)
}

type safeFetchedState struct {
	mu   sync.Mutex
	urls map[string]bool
}

var fetcheds safeFetchedState = safeFetchedState{
	urls: make(map[string]bool),
}

func updateFetchedState(newUrl string) {
	fetcheds.mu.Lock()
	fetcheds.urls[newUrl] = true
	fetcheds.mu.Unlock()
}

func isUrlAlreadyFetched(url string) bool {
	fetcheds.mu.Lock()
	defer fetcheds.mu.Unlock()

	_, is_fetched := fetcheds.urls[url]
	return is_fetched
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	if isUrlAlreadyFetched(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	updateFetchedState(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}
}

func WebCrawlerMain() {
	Crawl("https://golang.org/", 4, fetcher)
	time.Sleep(time.Second)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}

	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
