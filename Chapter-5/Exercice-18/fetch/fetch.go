// fetch is a program that fetches an URL and stores its contents into a local file
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// main is the entry point of the program
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <URL>\n", os.Args[0])
		os.Exit(0)
	}

	filename, length, err := fetch(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while fetching the URL: %v\n", err)
		os.Exit(2)
	}

	fmt.Printf("URL downloaded in local file %q (length=%d)\n", filename, length)
}

// fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, length int64, err error) {
	// Get the URL content
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	// Computes the local file name
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	// Create the local file
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		// Close file, but prefer error from Copy, if any.
		if closeErr := f.Close(); err == nil {
			// The returned value can be modified in the defer call
			err = closeErr
		}
	}()

	// Copy the URL content into the local file
	length, err = io.Copy(f, resp.Body)

	return local, length, err
}
