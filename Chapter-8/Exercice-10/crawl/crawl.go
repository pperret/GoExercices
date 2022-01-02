// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

// crawl encapsultes calls to Extract
func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

var cancel = make(chan struct{})

// main is the entry point of the program
func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	wg := sync.WaitGroup{}

	// Cancel crawling when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(cancel)
	loop:
		// Purge channel content
		for {
			select {
			case <-unseenLinks:
			default:
				break loop
			}
		}
	}()

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			for {
				select {
				case link := <-unseenLinks:
					foundLinks := crawl(link)
					go func() { worklist <- foundLinks }()
				case <-cancel:
					wg.Done()
					return
				}
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	wg.Add(1)
	go func() {
		for list := range worklist {
			for _, link := range list {
				select {
				case <-cancel:
					wg.Done()
					return
				default:
					if !seen[link] {
						seen[link] = true
						unseenLinks <- link
					}
				}
			}
		}
	}()

	// Wait for GO routines completion
	wg.Wait()
}
