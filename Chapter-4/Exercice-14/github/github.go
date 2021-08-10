// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import (
	"time"
)

// IssuesURL is the base URL to access github issues
const IssuesURL = "https://api.github.com/repos/golang/go/issues"

// IssuesListResult is the list of Issues
type IssuesListResult []Issue // Issues

// Issue is the data of an issue
type Issue struct {
	Number    int        `json:"number"`   // Issue ID
	HTMLURL   string     `json:"html_url"` // Url to access the issue
	Title     string     // Issue title
	State     string     // Issue state
	User      *User      // Issue creator
	Milestone *Milestone // Issue milestone
	CreatedAt time.Time  `json:"created_at"` // Issue creation timestamp
	Body      string     // Issue description (in Markdown format)
}

// User is the issue creator
type User struct {
	Login   string // Creator ID
	HTMLURL string `json:"html_url"` // URL to access to creator data
}

// Milestone is the milestone data
type Milestone struct {
	URL     string `json:"url"`      // URL to access to milestone data (JSON)
	HTMLURL string `json:"html_url"` // URL to access to milestone data (HTML)
	Title   string // Milestone title
	Number  int    `json:"number"` // Milestone ID
}
