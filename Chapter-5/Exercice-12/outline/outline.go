// outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// main is the entry point of the program
func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

// outline outlines the document tree of an HTML page identified by its URL
func outline(url string) error {

	// Get the URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Parse the HTML document
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	// depth is shared by both functions
	var depth int

	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}

	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}

	// Scan recursively the document tree
	forEachNode(doc, startElement, endElement)

	return nil
}

// forEachNode scans recursively the document tree
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
