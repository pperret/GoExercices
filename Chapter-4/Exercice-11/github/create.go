package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateIssue creates a new issue in the repository
func CreateIssue(username, token, owner, repo, title string) (*Issue, error) {
	// Build the URL
	postURL := strings.Join([]string{URLRepos, owner, repo, "issues"}, "/")

	// Get the body using the prefered editor
	body, err := getText()
	if err != nil {
		return nil, err
	}

	// Build the issue
	var issue Issue
	issue.Title = title
	issue.Body = body
	issue.State = "open"

	// Create the JSON request
	var request bytes.Buffer
	if err := json.NewEncoder(&request).Encode(&issue); err != nil {
		return nil, err
	}

	// Build the HTTP request
	client := &http.Client{}
	httpRequest, err := http.NewRequest("POST", postURL, &request)
	if err != nil {
		return nil, err
	}

	// Set authentication access
	httpRequest.SetBasicAuth(username, token)

	// Set content-type
	httpRequest.Header.Add("Content-Type", "application/json")

	// Set Accept header (recommanded by GitHub)
	httpRequest.Header.Add("Accept", "application/vnd.github.v3+json")

	// Send the request to GitHub
	resp, err := client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return nil, fmt.Errorf("create query failed: %s", resp.Status)
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
