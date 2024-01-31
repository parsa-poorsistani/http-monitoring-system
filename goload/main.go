package main

import (
	"fmt"
	"net/http"
	"sync"
)

func sendRequest() {
	url := "http://http-monitor-service.default.svc.cluster.local/api/server/all"

	for {
		_, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
		}
	}
}

func main() {
	// Number of goroutines to spawn
	numGoroutines := 100

	// Wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Start multiple goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sendRequest()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
