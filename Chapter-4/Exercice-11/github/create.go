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
	// Builds the URL
	postURL := strings.Join([]string{URLRepos, owner, repo, "issues"}, "/")

	// Gets the body using the prefered editor
	body, err := getText()
	if err != nil {
		return nil, err
	}

	// Builds the issue
	var issue Issue
	issue.Title = title
	issue.Body = body
	issue.State = "open"

	// Creates the JSON request
	var request bytes.Buffer
	if err := json.NewEncoder(&request).Encode(&issue); err != nil {
		return nil, err
	}

	// Builds the HTTP request
	client := &http.Client{}
	httpRequest, err := http.NewRequest("POST", postURL, &request)
	if err != nil {
		return nil, err
	}

	// Sets authentication access
	httpRequest.SetBasicAuth(username, token)

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
	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return nil, fmt.Errorf("create query failed: %s", resp.Status)
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
