package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type SafeVisited struct {
	mu      sync.Mutex
	visited map[string]bool
}

func (sv *SafeVisited) Check(url string) bool {
	sv.mu.Lock()
	defer sv.mu.Unlock()
	return sv.visited[url]
}

func (sv *SafeVisited) Mark(url string) {
	sv.mu.Lock()
	sv.visited[url] = true
	sv.mu.Unlock()
}

var sv SafeVisited = SafeVisited{
	visited: make(map[string]bool),
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, done chan string) {
	defer close(done)
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	if sv.Check(url) {
		return
	}
	sv.Mark(url)

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	res := make([]chan string, len(urls))
	for i, u := range urls {
		res[i] = make(chan string)
		go Crawl(u, depth-1, fetcher, res[i]) // I added the go part
	}

	// wait for all child goroutines to exit
	for i := range res {
		for s := range res[i] {
			done <- s
		}
	}
	return
}

func main() {
	done := make(chan string)
	Crawl("https://golang.org/", 4, fetcher, done)

	for s := range done {
		fmt.Println(s)
	}
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

// fetcher is a populated fakeFetcher.
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
