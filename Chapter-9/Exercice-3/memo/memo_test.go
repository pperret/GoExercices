package memo_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"

	"GoExercices/Chapter-9/Exercice-3/memo"
)

// httpRequestBody gets the content of a URL
func httpGetBody(url string, done <-chan struct{}) (interface{}, error) {
	// Create the request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Set the cancel channel of the request
	req.Cancel = done

	// Perform the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Get the content
	buf, err := ioutil.ReadAll(resp.Body)
	return buf, err
}

// incomingURLs returns a list of URLs to fetch in a channel
func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

// TestSequential accesses to URLs in a sequential way
func TestSequential(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, nil)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}

// TestConcurrent accesses to URLs in a parallel way
func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	var n sync.WaitGroup
	// Create the cancel channel
	done := make(chan struct{})
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url, done)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	//Cancel the pending requests by closing the channel
	close(done)
	n.Wait()
}
