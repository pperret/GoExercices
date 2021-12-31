// du computes the disk usage of the files in a directory.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

// fileStat is the data sent by go routines through the channel
type fileStat struct {
	rootIndex int   // Index of the root folder in the list
	size      int64 // File size
}

// rootStats is the statistics about a root folder
type rootStats struct {
	nfiles int64 // Files count
	nbytes int64 // Files size
}

// main is the entry point of the program
func main() {
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Channel to receive file stat sent by the go routines
	fileSizes := make(chan fileStat)

	// Statistics per root folder
	stats := make([]rootStats, len(roots))

	// Traverse each root of the file tree in parallel.
	var wg sync.WaitGroup
	for i, root := range roots {
		wg.Add(1)
		go walkDir(i, root, &wg, fileSizes)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}

loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			stats[size.rootIndex].nfiles++
			stats[size.rootIndex].nbytes += size.size
		case <-tick:
			printDiskUsage(roots, stats)
		}
	}

	printDiskUsage(roots, stats) // final totals
}

// printDiskUsage displays the current values of each root folder
func printDiskUsage(roots []string, stats []rootStats) {
	var nfiles, nbytes int64
	for i, root := range roots {
		nfiles += stats[i].nfiles
		nbytes += stats[i].nbytes
		fmt.Printf("%s: %d files  %.1f GB\n", root, stats[i].nfiles, float64(stats[i].nbytes)/1e9)
	}
	fmt.Printf("TOTAL: %d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(root int, dir string, wg *sync.WaitGroup, fileSizes chan<- fileStat) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(root, subdir, wg, fileSizes)
		} else {
			fileSizes <- fileStat{root, entry.Size()}
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
