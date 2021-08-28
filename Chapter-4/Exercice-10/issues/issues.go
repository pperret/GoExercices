// This program prints GitHub issues matching the search terms and sorted by date ranges
// (less than one month, less than one year, more than one year).
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"GoExercices/Chapter-4/Exercice-10/github"
)

func main() {

	// Get the issues according input parameters
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues in total\n", result.TotalCount)

	// Computes the milestone dates
	now := time.Now()
	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0, 0)

	// Creates the list of issues
	less1month := make([]*github.Issue, 0)
	less1year := make([]*github.Issue, 0)
	more1year := make([]*github.Issue, 0)

	// Sorts the issues
	for _, item := range result.Items {
		if item.CreatedAt.After(month) {
			less1month = append(less1month, item)
		} else if item.CreatedAt.After(year) {
			less1year = append(less1year, item)
		} else {
			more1year = append(more1year, item)
		}
	}

	// Displays the list of issues less than one month
	fmt.Printf("Less than one month (%d issues)\n", len(less1month))
	displayIssues(less1month)

	// Displays the list of issues less than one year
	fmt.Printf("Less than one year (%d issues)\n", len(less1year))
	displayIssues(less1year)

	// Displays the list of issues more than one year
	fmt.Printf("More than one year (%d issues)\n", len(more1year))
	displayIssues(more1year)
}

// Displays a list of issues
func displayIssues(issues []*github.Issue) {
	for _, issue := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
	}
}
