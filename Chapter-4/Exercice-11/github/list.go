package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ListIssues gets the issues list.
func ListIssues(owner, repos string) (*IssuesListResult, error) {
	// Builds the URL
	listURL := strings.Join([]string{URLRepos, owner, repos, "issues"}, "/")

	// Builds the HTTP request
	client := &http.Client{}
	httpRequest, err := http.NewRequest("GET", listURL, nil)
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
		return nil, fmt.Errorf("list query failed: %s", resp.Status)
	}

	// Decodes the response
	var result IssuesListResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
