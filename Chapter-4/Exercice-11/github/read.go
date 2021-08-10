package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ReadIssue reads issue data
func ReadIssue(owner string, repo string, issueNumber string) (*Issue, error) {
	// Create the request
	readURL := strings.Join([]string{"https://api.github.com/repos", owner, repo, "issues", issueNumber}, "/")

	// Send the request to github
	resp, err := http.Get(readURL)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Read query failed: %s", resp.Status)
	}

	// Decode the response
	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}
