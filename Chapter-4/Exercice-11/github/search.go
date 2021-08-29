package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	// Builds the URL
	searchURL := URLSearch + "?q=" + url.QueryEscape(strings.Join(terms, " "))

	// Builds the HTTP request
	client := &http.Client{}
	httpRequest, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return nil, err
	}

	// Sets content-type
	httpRequest.Header.Add("Content-Type", "application/json")

	// Sets Accept header (recommanded by GitHub)
	httpRequest.Header.Add("Accept", "application/vnd.github.v3+json")

	// Sends the request to GitHub
	resp, err := client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	// Decodes the response
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
