package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

const postTimeout = 2 * time.Second

// Custom client to enable canceling the request via the transport
var transport = http.Transport{}

var client = http.Client{
	Transport: &transport,
}

// TODO: Implement
func getWebhooks() []string {
	return []string{}
}

// Triggers the handlers associated with the hook
func trigger(payload interface{}) (int, error) {
	urls := getWebhooks()

	n := len(urls)

	if n == 0 {
		return 0, nil
	}

	// Encode the data as JSON for all hook handlers to receive
	var data bytes.Buffer

	encoder := json.NewEncoder(&data)

	if err := encoder.Encode(payload); err != nil {
		return 0, err
	}

	wg := sync.WaitGroup{}
	defer wg.Wait()

	wg.Add(n)

	for _, url := range urls {
		go func(url string) {
			post(url, &data, postTimeout)
			wg.Done()
		}(url)
	}

	return n, nil
}

// Send a POST request to the URL
func post(url string, data io.Reader, timeout time.Duration) {
	req, _ := http.NewRequest("POST", url, data)

	timer := time.AfterFunc(timeout, func() {
		transport.CancelRequest(req)
		log.Println(url, "timed out")
	})

	defer timer.Stop()

	client.Do(req)
}
