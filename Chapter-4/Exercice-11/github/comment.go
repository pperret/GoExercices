package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// AddComment adds a comment to an existing issue
func AddComment(username, token, owner, repo, issueNumber string) (*Issue, error) {
	// Build the URL
	postURL := strings.Join([]string{URLRepos, owner, repo, "issues", issueNumber, "comments"}, "/")

	// Get the comment using the prefered editor
	text, err := getText()
	if err != nil {
		return nil, err
	}

	// Build the issue
	var comment Comment
	comment.Body = text

	// Create the JSON request
	var request bytes.Buffer
	if err := json.NewEncoder(&request).Encode(&comment); err != nil {
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

	// Send the request to github
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
