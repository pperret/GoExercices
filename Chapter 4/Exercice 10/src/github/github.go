// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import (
	"time"
)

// IssuesURL is the base URL to access github issues
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult is the list of Issues provided by search
type IssuesSearchResult struct {
	TotalCount int      `json:"total_count"` // Number of issues
	Items      []*Issue // Issues
}

// Issue is the data of an issue
type Issue struct {
	Number    int       // Issue ID
	HTMLURL   string    `json:"html_url"` // Url to access the issue
	Title     string    // Issue title
	State     string    // Issue state
	User      *User     // Issue creator
	CreatedAt time.Time `json:"created_at"` // Issue creation timestamp
	Body      string    // Issue description (in Markdown format)
}

// User is the issue creator
type User struct {
	Login   string // Creator ID
	HTMLURL string `json:"html_url"` // URL to access to creator data
}
