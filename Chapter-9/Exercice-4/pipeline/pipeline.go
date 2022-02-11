// pipeline is an example of a pipeline built using go routines
// On my Mac, I reached 90.000.000 nodes and the propagation time was 17m 48s.
package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// routine is the implement of a stage of the pipeline
func routine(in <-chan int, out chan<- int) {
	for i := range in {
		out <- i
	}
}

// build builds count stages of pipeline
func build(count int, wg sync.WaitGroup) (chan<- int, <-chan int) {
	in := make(chan int)
	out := in
	for i := 0; i < count; i++ {
		wg.Add(1)
		tmp_in := out
		tmp_out := make(chan int)
		go func(in <-chan int, out chan<- int) {
			defer wg.Done()
			routine(in, out)
		}(tmp_in, tmp_out)
		out = tmp_out
	}
	return in, out
}

// test returns the duration to send an integer through the pipeline
func test(in chan<- int, out <-chan int) (time.Duration, error) {

	now := time.Now()
	in <- 10
	v := <-out
	duration := time.Since(now)

	if v != 10 {
		return 0, fmt.Errorf("invalid value sent through the pipe")
	}
	return duration, nil
}

// usage displays the CLI interface
func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <stage count>\n", os.Args[0])
}

// main is the entry point of the program
func main() {

	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
		os.Exit(1)
	}

	var wg sync.WaitGroup
	in, out := build(count, wg)

	fmt.Printf("Pipeline is built.\n")

	duration, err := test(in, out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(2)
	}

	fmt.Printf("Pipeline is closing...\n")

	close(in)
	wg.Wait()
	fmt.Printf("Count:%d, Duration: %v\n", count, duration)
}
