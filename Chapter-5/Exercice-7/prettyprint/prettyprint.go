// PrettyPrint formats a HTML document
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// main is the entry point of the program
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <url>\n", os.Args[0])
		os.Exit(1)
	}

	err := prettyPrint(os.Stdout, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error managing URL: %v\n", err)
		os.Exit(2)
	}
}

// prettyPrint downloads then formats the HTML document.
func prettyPrint(w io.Writer, url string) (err error) {

	// Get the HTML document
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	// Parse the document
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("error while parsing HTML: %s", err)
		return
	}

	// Format the document
	forEachNode(w, doc, startElement, endElement)
	return
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(w io.Writer, n *html.Node, pre, post func(w io.Writer, n *html.Node)) {
	if pre != nil {
		pre(w, n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(w, c, pre, post)
	}
	if post != nil {
		post(w, n)
	}
}

var depth int

// startElement is called before managing childs
func startElement(w io.Writer, n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		fmt.Fprintf(w, "%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Fprintf(w, " %s='%s'", a.Key, a.Val)
		}
		if n.FirstChild == nil {
			fmt.Fprintf(w, "/>\n")
		} else {
			fmt.Fprintf(w, ">\n")
		}
		depth++

	case html.CommentNode:
		fmt.Fprintf(w, "%*s<!-- %s -->\n", depth*2, "", n.Data)

	case html.TextNode:
		t := strings.TrimSpace(n.Data)
		if len(t) > 0 {
			lines := strings.Split(t, "\n")
			for _, l := range lines {
				fmt.Fprintf(w, "%*s%s\n", depth*2, "", l)
			}
		}
	}
}

// endElement is called after managing childs
func endElement(w io.Writer, n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
