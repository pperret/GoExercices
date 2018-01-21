// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	// Create one channel to get goroutine results
	ch := make(chan string)

	// Launch goroutines for each URL
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	// Get results from goroutines
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// Fetch one URL
func fetch(url string, ch chan<- string) {
	start := time.Now()

	// Access to the URL
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	// Compute the output file name
	name := strings.Split(url, "//")[1] + ".html"

	// Open the output file
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		ch <- fmt.Sprintf("while openning %s: %v", name, err)
		return
	}

	// Store URL content into output file
	nbytes, err := io.Copy(f, resp.Body)

	// Close URL response and output file
	resp.Body.Close() // don't leak resources
	f.Close()

	// Check storing result
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()

	// Send result to main
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
