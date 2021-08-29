package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ReadIssue reads issue data
func ReadIssue(owner string, repo string, issueNumber string) (*Issue, error) {
	// Builds the URL
	readURL := strings.Join([]string{URLRepos, owner, repo, "issues", issueNumber}, "/")

	// Builds the HTTP request
	client := &http.Client{}
	httpRequest, err := http.NewRequest("GET", readURL, nil)
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
		return nil, fmt.Errorf("read query failed: %s", resp.Status)
	}

	// Decodes the response
	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}
