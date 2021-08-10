// findlinks crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"golang.org/x/net/html"
)

// main is the entry point of the program
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				links := f(item)
				if len(links) > 0 {
					parsedURL, err := url.Parse(item)
					if err != nil {
						log.Print(err)
					}
					// Only link of current domain are considered
					for _, link := range links {
						parsedLink, err := url.Parse(link)
						if err != nil {
							log.Print(err)
							continue
						}
						if parsedLink.Scheme != parsedURL.Scheme || len(parsedLink.Scheme) == 0 {
							continue
						}
						if parsedLink.Host != parsedURL.Host || len(parsedLink.Host) == 0 {
							continue
						}
						worklist = append(worklist, link)
					}
				}
			}
		}
	}
}

// crawl gets the list of links in the HTML page specified by url parameter
func crawl(url string) []string {
	fmt.Println(url)
	// Save the content
	err := saveContent(url)
	if err != nil {
		log.Print(err)
		return nil
	}

	// Extract the links from the URL
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// Extract makes an HTTP GET request to the specified URL,
// parses the response as HTML, and returns
// the links in the HTML document.
func Extract(url string) ([]string, error) {
	// Get the URL content
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	// Get the links in the HTML page
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// saveContent saves the content of the URL
func saveContent(htmlURL string) error {
	// Get the URL content
	resp, err := http.Get(htmlURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getting %s: %s", htmlURL, resp.Status)
	}

	// Parse the URL
	parsedURL, err := url.Parse(htmlURL)
	if err != nil {
		return err
	}

	// Build the file name
	path := filepath.FromSlash(parsedURL.Path)
	if len(path) == 0 {
		path = "default.html"
	} else {
		if path[len(path)-1] == os.PathSeparator {
			path = filepath.Join(path, "default.html")
		}
	}
	path = filepath.Join(parsedURL.Host, path)

	// Create the folder
	dir := filepath.Dir(path)
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}

	// Save the UTL content
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
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
