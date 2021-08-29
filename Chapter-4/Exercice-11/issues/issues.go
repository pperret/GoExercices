// Issues manages GitHub issues
package main

import (
	"fmt"
	"log"
	"os"

	"GoExercices/Chapter-4/Exercice-11/github"
)

// main is the entry point of the program
func main() {

	// Checks parameters
	if len(os.Args) < 2 {
		usage(os.Args[0])
		os.Exit(1)
	}

	switch os.Args[1] {
	// Lists issues
	case "list":
		if len(os.Args) < 4 {
			fmt.Fprintf(os.Stderr, "Usage: %s list <owner> <repository>\n", os.Args[0])
			os.Exit(1)
		}
		result, err := github.ListIssues(os.Args[2], os.Args[3])
		if err != nil {
			log.Fatal(err)
		}
		for _, item := range *result {
			fmt.Printf("#%5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}

	// Searchs issues
	case "search":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "Usage: %s search <param> [params...]\n", os.Args[0])
			os.Exit(1)
		}
		// Gets the issue according on input parameters
		result, err := github.SearchIssues(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%d issues:\n", result.TotalCount)
		for _, item := range result.Items {
			fmt.Printf("#%5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}

	// Creates an issue
	case "create":
		if len(os.Args) < 7 {
			fmt.Fprintf(os.Stderr, "Usage: %s create <user|email> <token> <owner> <repository> <title>\n", os.Args[0])
			os.Exit(1)
		}
		result, err := github.CreateIssue(os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("#%5d %9.9s %.55s\n", result.Number, result.User.Login, result.Title)

	// Reads a single issue
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

	// Appends a comment to an existing issue
	case "comment":
		if len(os.Args) < 7 {
			fmt.Fprintf(os.Stderr, "Usage: %s comment <user|email> <token> <owner> <repository> <issue number>\n", os.Args[0])
			os.Exit(1)
		}
		_, err := github.AddComment(os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6])
		if err != nil {
			log.Fatal(err)
		}

	// Closes an issue (an issue cannot be deleted)
	case "close":
		if len(os.Args) < 7 {
			fmt.Fprintf(os.Stderr, "Usage: %s close <user|email> <token> <owner> <repository> <issue number>\n", os.Args[0])
			os.Exit(1)
		}
		_, err := github.CloseIssue(os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6])
		if err != nil {
			log.Fatal(err)
		}
	default:
		usage(os.Args[0])
		os.Exit(2)
	}
}

func usage(argv0 string) {
	fmt.Fprintf(os.Stderr, "Usage: %s <cmd> [args]\n", argv0)
	fmt.Fprintf(os.Stderr, "Available commands are: 'list', 'search', 'create', 'read', 'comment' and 'close'\n")
}
