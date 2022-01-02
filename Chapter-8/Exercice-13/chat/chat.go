// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const idleTimeout = 5 * time.Minute

type client chan<- string // outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

// broadcaster sends messages to each connected client
func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for client := range clients {
				client <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

// handleConn manages a connection with a user
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	timer := time.NewTimer(idleTimeout)
	go clientIdle(conn, timer)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		timer.Reset(idleTimeout)
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	timer.Stop()
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

// clientWriter is a go routine to send message to a user
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

// clientIdle detects client inactivity
func clientIdle(conn net.Conn, timer *time.Timer) {
	<-timer.C
	conn.Close()
}

// main is the entry point of the program
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
