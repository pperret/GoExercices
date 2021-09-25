// Findlinks prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// main is the entry point of the program
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse HML document: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// Map containing the link's attribute name for each targeted element
var linkAttributes = map[string][]string{
	"a":      {"href"},
	"link":   {"href"},
	"iframe": {"src"},
	"script": {"src"},
	"img":    {"src"},
	"form":   {"action"},
	"html":   {"manifest"},
	"video":  {"src", "poster"}}

// visit appends to links each link found in node and returns the result.
func visit(links []string, node *html.Node) []string {

	// Look for a link in the current element
	if node.Type == html.ElementNode {
		attributeNames, ok := linkAttributes[node.Data]
		if ok {
			// Search for the corresponding attribute
			for _, attributeName := range attributeNames {
				for _, a := range node.Attr {
					if a.Key == attributeName {
						links = append(links, a.Val)
					}
				}
			}
		}
	}

	// Visit the childs (recursively)
	if node.FirstChild != nil {
		links = visit(links, node.FirstChild)
	}

	// Visit the siblings (recursively)
	if node.NextSibling != nil {
		links = visit(links, node.NextSibling)
	}

	return links
}
