// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"GoExercices/Chapter-4/Exercice-10/github"
)

func main() {

	// Get the issue according input parameters
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues in total\n", result.TotalCount)

	now := time.Now()
	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0, 0)

	less1month := make([]*github.Issue, 0)
	less1year := make([]*github.Issue, 0)
	more1year := make([]*github.Issue, 0)

	// Sort the issues
	for _, item := range result.Items {
		if item.CreatedAt.After(month) {
			less1month = append(less1month, item)
		} else if item.CreatedAt.After(year) {
			less1year = append(less1year, item)
		} else {
			more1year = append(more1year, item)
		}
	}

	// Less than one month
	fmt.Printf("Less than one month (%d issues)\n", len(less1month))
	for _, item := range less1month {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

	// Less than one year
	fmt.Printf("Less than one year (%d issues)\n", len(less1year))
	for _, item := range less1year {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

	// More than one year
	fmt.Printf("More than one year (%d issues)\n", len(more1year))
	for _, item := range more1year {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
