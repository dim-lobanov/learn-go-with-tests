package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			//results[u] = wc(u)
			resultChannel <- result{u, wc(u)} // Send statement
		}(url)
		// func(){}() <- func definition func(){} + call ()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel // Receive expression
		results[r.string] = r.bool
	}

	return results
}

// Because the only way to start a goroutine is to put go in front of a function call,
// we often use anonymous functions when we want to start a goroutine.

// there is race detector in Go: go test -race
// https://go.dev/blog/race-detector
