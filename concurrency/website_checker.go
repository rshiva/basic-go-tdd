package concurrency

type WebsiteChecker func(string) bool

// As we don't need either value to be named, each of them is anonymous within the struct;
//  this can be useful in when it's hard to know what to name a value.
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// Send statement
			resultsChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// Receive expression
		result := <-resultsChannel
		results[result.string] = result.bool
	}
	// time.Sleep(2 * time.Second)
	return results
}
