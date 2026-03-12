package checker

import (
	"net/http"
	"sync"
	"time"
)

// CheckURL checks a single URL and sends the result on ch.
func CheckURL(url string, ch chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- Result{URL: url, Status: "Down", Latency: 0}
		return
	}
	defer resp.Body.Close()

	ch <- Result{
		URL:     url,
		Status:  resp.Status,
		Latency: time.Since(start),
	}
}

// Run checks all URLs concurrently and returns results in order received.
// It blocks until all checks complete.
func Run(urls []string) []Result {
	if len(urls) == 0 {
		return nil
	}

	resultsCh := make(chan Result, len(urls))
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go CheckURL(url, resultsCh, &wg)
	}

	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	var results []Result
	for res := range resultsCh {
		results = append(results, res)
	}
	return results
}
