// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import (
	"time"
)

// Base URL to access github issues
const IssuesURL = "https://api.github.com/search/issues"

// Search result
type IssuesSearchResult struct {	
	TotalCount int `json:"total_count"`		// Number of issues
	Items      []*Issue						// Issues
}

// Issue
type Issue struct {	
	Number    int								// Issue ID	
	HTMLURL   string `json:"html_url"`			// Url to access the issue	
	Title     string `json:"title,omitempty"`	// Issue title
	State     string `json:"state"`				// Issue state
	User      *User								// Issue creator
	CreatedAt time.Time `json:"created_at"`		// Issue creation timestamp
	Body      string `json:"body"`				// Issue description (in Markdown format)
}

// Issue creator
type User struct {
	Login   string						// Creator ID
	HTMLURL string `json:"html_url"`	// URL to access to creator data
}

// Issue comment
type Comment struct {
	Body      string `json:"body"`	// Comment description (in Markdown format)	
}
