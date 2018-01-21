// Issues manages GitHub issues
package main

import (
	"fmt"
	"github"
	"log"
	"os"
)

// main is the entry point of the program
func main() {

	// Check parameters
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <cmd> [args]\n", os.Args[0])
		os.Exit(1)
	}

	switch os.Args[1] {
	// Read a issue list
	case "list":
		// Get the issue according input parameters
		result, err := github.SearchIssues(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%d issues:\n", result.TotalCount)
		for _, item := range result.Items {
			fmt.Printf("#%5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}

	// Create an issue
	case "create":
		if len(os.Args) < 7 {
			fmt.Fprintf(os.Stderr, "Usage: %s create <user> <password> <owner> <repository> <title>\n", os.Args[0])
			os.Exit(1)
		}
		result, err := github.CreateIssue(os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("#%5d %9.9s %.55s\n", result.Number, result.User.Login, result.Title)

	// Read a single issue
	case "read":
		if len(os.Args) < 5 {
			fmt.Fprintf(os.Stderr, "Usage: %s read <owner> <repository> <issue number>\n", os.Args[0])
			os.Exit(1)
		}
		result, err := github.ReadIssue(os.Args[2], os.Args[3], os.Args[4])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("#%5d %9.9s %.55s\n", result.Number, result.User.Login, result.Title)

	// Append a comment to an existing issue
	case "comment":
		if len(os.Args) < 7 {
			fmt.Fprintf(os.Stderr, "Usage: %s comment <user> <password> <owner> <repository> <issue number>\n", os.Args[0])
			os.Exit(1)
		}
		_, err := github.AddComment(os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6])
		if err != nil {
			log.Fatal(err)
		}

	// Close an issue (an issue cannot be deleted)
	case "close":
		if len(os.Args) < 7 {
			fmt.Fprintf(os.Stderr, "Usage: %s close <user> <password> <owner> <repository> <issue number>\n", os.Args[0])
			os.Exit(1)
		}
		_, err := github.CloseIssue(os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6])
		if err != nil {
			log.Fatal(err)
		}
	}
}
