// pingpong is a sample of goroutines exchanging messages through channels
package main

import (
	"fmt"
	"sync"
	"time"
)

// MAX is the test duration (in seconds)
const MAX = 10

// exchange receives an integer and sends back the value +1
// Counter is incremented
func exchange(counter *uint64, in <-chan uint64, out chan<- uint64, cancel <-chan bool) {
loop:
	for {
		select {
		case <-in:
			*counter++
			out <- *counter
		case <-cancel:
			<-in // Consume the data to prevent deadlook
			break loop
		}
	}
	close(out)
}

// main is the program entry point
func main() {
	ch1 := make(chan uint64)
	ch2 := make(chan uint64)
	cancel := make(chan bool)

	var wg sync.WaitGroup

	wg.Add(2)
	var c1, c2 uint64
	go func() {
		defer wg.Done()
		exchange(&c1, ch1, ch2, cancel)
	}()
	go func() {
		defer wg.Done()
		exchange(&c2, ch2, ch1, cancel)
	}()

	// Start the ping-pong
	ch1 <- 0

	// Delay for test duration
	time.Sleep(MAX * time.Second)

	// Stop the goroutines
	close(cancel)

	// Wait for goutines completion
	wg.Wait()

	// Display the result
	fmt.Printf("count: %d, throughput: %f ping-pong/s\n", c1, float64(c1)/MAX)
}
