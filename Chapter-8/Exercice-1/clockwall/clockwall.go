package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type Server struct {
	Name string
	Addr string
	Time string
}

// getServerList returns the list of servers privided as arguments of the program
func getServerList() ([]*Server, error) {
	var serverList []*Server
	for _, server := range os.Args[1:] {
		fields := strings.Split(server, "=")
		if len(fields) != 2 {
			return nil, fmt.Errorf("invalid argument format: %s", server)
		}

		srv := Server{Name: fields[0], Addr: fields[1]}
		serverList = append(serverList, &srv)
	}
	return serverList, nil
}

// getTime gets the time of a remote server
func getTime(addr string) (string, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return "", fmt.Errorf("unable to connect to remote server: %v", err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("unable to read time: %v", err)
	}
	return str[:len(str)-1], nil
}

// poolServerTime updates the time in the server's structure periodically
func poolServerTime(srv *Server) {
	for {
		t, err := getTime(srv.Addr)
		if err != nil {
			srv.Time = "---"
			continue
		} else {
			srv.Time = t
		}

		time.Sleep(100 * time.Millisecond)
	}
}

// main is the entry point of the program
func main() {

	// Check that at least one server is provided
	if len(os.Args[1:]) == 0 {
		fmt.Fprintf(os.Stderr, "No enough argument\n")
		os.Exit(1)
	}

	// Get the server list
	serverList, err := getServerList()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get server list: %v\n", err)
		os.Exit(2)
	}

	// Start Go routines (one per server) to update sever's time in background
	for _, srv := range serverList {
		go poolServerTime(srv)
	}

	// Wait during the launch of go routines
	time.Sleep(100 * time.Millisecond)

	// Display the current time of each server
	for {
		fmt.Printf("%-40s%-30s\n", "SERVER", "TIME")
		for _, srv := range serverList {
			if srv.Time != "" {
				fmt.Printf("%-40s%-30s\n", srv.Name, srv.Time)
			}
		}
		fmt.Println()

		time.Sleep(1 * time.Second)
	}
}
