// CountWords prints the numbers of words and images of an HTML document downloaded from the Internet.
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// main is the entry point of the program
func main() {
	// Loop on each URL
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while reading URL: %v\n", err)
			continue
		}
		fmt.Printf("%s: %d words, %d images\n", url, words, images)
	}
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {

	// Get the HTML document
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	// Parse the document
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("Error while parsing HTML: %s", err)
		return
	}

	// Count words and images
	words, images = countWordsAndImages(doc)
	return
}

// countWordsAndImages counts words and images in a HTML document
func countWordsAndImages(node *html.Node) (words, images int) {
	// Count number of words in a text node
	if node.Type == html.TextNode {
		t := strings.TrimSpace(node.Data)
		if len(t) != 0 {
			input := bufio.NewScanner(strings.NewReader(t))
			input.Split(bufio.ScanWords)
			for input.Scan() {
				input.Text()
				words++
			}
		}
	}

	// Detect image
	if node.Type == html.ElementNode {
		if node.Data == "img" {
			images++
		}
	}

	// Recurse only for none script or style elements
	if node.Type != html.ElementNode || (node.Data != "script" && node.Data != "style") {
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			w, i := countWordsAndImages(child)
			words += w
			images += i
		}
	}
	return
}
