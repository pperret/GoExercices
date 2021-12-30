// Reverb is a TCP server that simulates an echo.
// clients are automatically disconnected after 10 seconds of inactivity
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// echo returns the received text three times
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

// handleConn manages incoming connections
func handleConn(c net.Conn) {
	defer c.Close()
	texts := make(chan string)
	go func() {
		input := bufio.NewScanner(c)
		// NOTE: ignoring potential errors from input.Err()
		for input.Scan() {
			texts <- input.Text()
		}
	}()

	timer := time.NewTimer(10 * time.Second)
	for {
		select {
		case text := <-texts:
			timer.Reset(10 * time.Second)
			go func() {
				echo(c, text, 1*time.Second)
			}()
		case <-timer.C:
			// The connection is closed. Therefore, the input go routine will exit
			// WaitGroup is not useful because echo duration is only 2 seconds
			return
		}
	}
}

// main is the entry point of the program
func main() {

	// Create the listen endpoint
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	// Manage incoming connections
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
