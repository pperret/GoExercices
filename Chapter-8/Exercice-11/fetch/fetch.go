// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// main is the entry point of the program
func main() {

	// At least one URL is required
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <URL>...\n", os.Args[0])
		os.Exit(1)
	}

	// Fetch the URL
	b, err := mirroredQuery(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetch error: %v\n", err)
		os.Exit(2)
	}

	// Print the body
	fmt.Println(b)
}

// Response contains the reply of a go routine fetching one of the requested URL
type Response struct {
	body string
	err  error
}

// mirroredQuery performs URL fetching in parallel and returns the quickest reply
func mirroredQuery(urls []string) (string, error) {
	done := make(chan struct{})
	responses := make(chan Response)
	for _, url := range urls {
		go func(url string) {
			response, err := fetch(url, done)
			if err != nil {
				responses <- Response{"", err}
			} else {
				close(done)                          // Cancel other requests
				responses <- Response{response, nil} // Send the reply to the output channel
			}
		}(url)
	}

	// Return the quickest valid response
	for range urls {
		resp := <-responses
		if resp.err == nil {
			return resp.body, nil
		}
	}
	return "", fmt.Errorf("no valid reponse")
}

// fetch fetches one URL with a cancel channel
func fetch(url string, done chan struct{}) (string, error) {
	// Create the request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	// Set the cancel channel of the request
	req.Cancel = done

	// Perform the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
