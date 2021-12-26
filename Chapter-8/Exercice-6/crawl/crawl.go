// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Work is a list of URLs returned by the crawler
type Work struct {
	urls  []string
	depth int
}

// Link is a URL to analyze by the crawler
type Link struct {
	url   string
	depth int
}

// crawl extracts links from a document specified by its URL
func crawl(maxDepth int, link *Link) []string {
	fmt.Println(link.url)
	if link.depth < maxDepth {
		list, err := Extract(link.url)
		if err != nil {
			log.Print(err)
			return nil
		}
		return list
	}
	return nil
}

// main is the entry point of the program
func main() {

	// Get depth limit as an optional flag
	maxDepth := flag.Int("depth", 3, "Depth limit")
	flag.Parse()

	//  Check the URLs list is not empty
	if len(flag.Args()) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s <URL>...\n", os.Args[0])
		os.Exit(1)
	}

	workList := make(chan Work)    // lists of URLs, may have duplicates
	unseenLinks := make(chan Link) // de-duplicated URLs

	// Add command-line arguments (URLs) to worklist.
	go func() { workList <- Work{flag.Args(), 0} }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(*maxDepth, &link)
				currentDepth := link.depth
				if foundLinks != nil {
					go func() { workList <- Work{foundLinks, currentDepth + 1} }()
				}
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for work := range workList {
		for _, url := range work.urls {
			if !seen[url] {
				seen[url] = true
				unseenLinks <- Link{url, work.depth}
			}
		}
	}
}
