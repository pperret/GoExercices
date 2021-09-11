// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import (
	"time"
)

// URLSearch is the base URL to search in GitHub issues
const URLSearch = "https://api.github.com/search/issues"

// URLRepos is the base URL to access to the GitHub repository
const URLRepos = "https://api.github.com/repos"

// IssuesSearchResult is the list of Issues resulting of a search operation
type IssuesSearchResult struct {
	TotalCount int      `json:"total_count"` // Number of issues
	Items      []*Issue // Issues
}

// IssuesListResult is the list of GitHub issues
type IssuesListResult []Issue

// Issue is the data of a GitHub issue
type Issue struct {
	Number    int       // Issue ID
	HTMLURL   string    `json:"html_url"`        // Url to access the issue
	Title     string    `json:"title,omitempty"` // Issue title
	State     string    `json:"state"`           // Issue state
	User      *User     // Issue creator
	CreatedAt time.Time `json:"created_at"` // Issue creation timestamp
	Body      string    `json:"body"`       // Issue description (in Markdown format)
}

// User is a GitHub issue creator
type User struct {
	Login   string // Creator ID
	HTMLURL string `json:"html_url"` // URL to access to creator data
}

// Comment is a GitHub issue comment
type Comment struct {
	Body string `json:"body"` // Comment description (in Markdown format)
}
