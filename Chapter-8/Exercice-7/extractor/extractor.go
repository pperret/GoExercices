// extractor creates a disk copy of a web site
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

// main is the entry point of the program
func main() {

	//  Check the arguments
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <URL> <folder>\n", os.Args[0])
		os.Exit(1)
	}

	baseURL := os.Args[1]
	targetFolder := os.Args[2]

	workList := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add the base URL to the worklist.
	go func() { workList <- []string{baseURL} }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link, baseURL, targetFolder)
				if foundLinks != nil {
					go func() { workList <- foundLinks }()
				}
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for work := range workList {
		for _, url := range work {
			if !seen[url] {
				seen[url] = true
				unseenLinks <- url
			}
		}
	}
}

// crawl extracts a document specified by its URL and returns its internal local links
func crawl(link string, baseURL string, targetFolder string) []string {
	list, err := Extract(link, baseURL, targetFolder)
	if err != nil {
		log.Printf("Unable to extract %q: %v\n", link, err)
		return nil
	}
	return list
}

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the local links in the HTML document.
func Extract(documentURL string, baseURL string, targetFolder string) ([]string, error) {

	// Get the UTL content
	resp, err := http.Get(documentURL)
	if err != nil {
		return nil, fmt.Errorf("unable to download document %q: %v", documentURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error while getting %q: %s", documentURL, resp.Status)
	}

	// Extract content according to content type
	if !IsHtml(resp) {
		err = ExtractBinary(targetFolder, resp)
		if err != nil {
			return nil, fmt.Errorf("unable to save resource %q: %v", documentURL, err)
		}
		return []string{}, nil
	} else {
		links, err := ExtractHTML(documentURL, baseURL, targetFolder, resp)
		if err != nil {
			return nil, fmt.Errorf("unable to save document %q: %v", documentURL, err)
		}
		return links, nil
	}
}

// IsHtml determines is the response is a HTML document
func IsHtml(resp *http.Response) bool {
	for _, ct := range resp.Header["Content-Type"] {
		parts := strings.Split(ct, ";")
		part := strings.TrimSpace(parts[0])
		if part == "text/html" {
			return true
		}
	}
	return false
}

// ExtractBinary saves a HTML document
func ExtractHTML(documentURL string, baseURL string, targetFolder string, resp *http.Response) ([]string, error) {
	// Parse the HTML document
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while parsing %q as HTML: %v", documentURL, err)
	}

	// Extract local links using the forEachNode function
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "link") {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				linkString := link.String()
				if !strings.HasPrefix(linkString, baseURL) {
					continue // ignore external URLs
				}

				// Append the link to the list
				links = append(links, linkString)

				// Update local target link
				a.Val = buildTargetLink(link)
			}
		}
	}
	forEachNode(doc, visitNode, nil)

	// Generate the target file name
	fileName := buildTargetFileName(targetFolder, resp.Request.URL)

	// Create the target folder
	err = os.MkdirAll(filepath.Dir(fileName), os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("unable to create target folder %q: %v", filepath.Dir(fileName), err)
	}

	// Create the target file
	f, err := os.Create(fileName)
	if err != nil {
		return nil, fmt.Errorf("unable to create target file %q: %v", fileName, err)
	}
	defer f.Close()

	// Save the document
	err = html.Render(f, doc)
	if err != nil {
		return nil, fmt.Errorf("unable to write into target file %q: %v", fileName, err)
	}
	return links, nil
}

// ExtractBinary saves a non-HTML resource
func ExtractBinary(targetFolder string, resp *http.Response) error {
	// Generate the target file name
	fileName := buildTargetFileName(targetFolder, resp.Request.URL)

	// Create the target folder
	err := os.MkdirAll(filepath.Dir(fileName), os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to create target folder %q: %v", filepath.Dir(fileName), err)
	}

	// Create the target file
	f, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("unable to create target file %q: %v", fileName, err)
	}
	defer f.Close()

	// Save the resource
	io.Copy(f, resp.Body)
	return nil
}

// HasKnownSuffix determines is the path extension is known
func HasKnownSuffix(path string) bool {
	var knownSuffixes = [...]string{".html", ".htm", ".css"}
	for _, s := range knownSuffixes {
		if strings.HasSuffix(path, s) {
			return true
		}
	}
	return false
}

// buildTargetFileName generates the local full name of a downloaded document
func buildTargetFileName(targetFolder string, targetURL *url.URL) string {
	var fileName string
	if targetURL.Path == "" {
		fileName = path.Join(targetFolder, targetURL.Host, "index.html")
	} else if strings.HasSuffix(targetURL.Path, "/") {
		fileName = path.Join(targetFolder, targetURL.Host, targetURL.Path, "index.html")
	} else if !HasKnownSuffix(targetURL.Path) {
		fileName = path.Join(targetFolder, targetURL.Host, targetURL.Path, "index.html")
	} else {
		fileName = path.Join(targetFolder, targetURL.Host, targetURL.Path)
	}
	return fileName
}

// buildTargetLink generates a local target link for a referenced document
func buildTargetLink(targetURL *url.URL) string {
	var targetLink string
	if targetURL.Path == "" {
		targetLink = path.Join("index.html")
	} else if strings.HasSuffix(targetURL.Path, "/") {
		targetLink = path.Join(".", targetURL.Path, "index.html")
	} else if !HasKnownSuffix(targetURL.Path) {
		targetLink = path.Join(".", targetURL.Path, "index.html")
	} else {
		targetLink = path.Join(".", targetURL.Path)
	}
	return "file://" + targetLink

}

// forEachNode applies pre and post functions to each node
func forEachNode(node *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(node)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(node)
	}
}
